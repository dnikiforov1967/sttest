package asyncservice

import "net/http"
import "encoding/json"
import "strconv"
import "time"
import "math"
import "github.com/gorilla/mux"
import "../errhand"
import "../config"
import "../access"

func initiateTaskMap() map[uint64]*TaskResponse {
	tempRef := mapAccess.Load()
	if tempRef!=nil {
		return *tempRef.(*map[uint64]*TaskResponse)
	} else {
		    mapLock.Lock()
			defer mapLock.Unlock()
			tempRef = mapAccess.Load()
			if (tempRef != nil) {
				return *tempRef.(*map[uint64]*TaskResponse)
			} else {
				taskMap := make(map[uint64]*TaskResponse)
				mapAccess.Store(&taskMap)
				return taskMap
			}
	}
}

func proceed(id uint64, isin string, underlying float64, volatility float64, signalChan chan int) {
	respMap := initiateTaskMap();
	response := TaskResponse{id, isin, StatusInProgress, 0, ""}
	respMap[id] = &response;

	initTime := time.Now()
	
	//Summary time ~ 5 sec
	for i := 0; i<10; i++ {
		//Here we simulate steps of price calculation
		timer := time.NewTimer(500 * time.Millisecond)
		<- timer.C
		if checkTimeOut(&initTime) {
			response.Status = StatusTimedOut
			if signalChan != nil {
				signalChan <- -1
			}
			break;
		}
	}
	//Normal commitment
	response.Status = StatusCompleted
	response.Price = math.Round(underlying*volatility*100000)/100
	response.PriceDate = time.Now().Format(time.RFC3339)
	if signalChan != nil {
		signalChan <- 0
	}
}

func checkTimeOut(initTime *time.Time) bool {
	duration := time.Since(*initTime)
	var millisec int64 = duration.Nanoseconds()/1000000
	if millisec >= config.TimeOut {
		return true
	}
	return false
}

func getTaskState(id uint64) (TaskResponse, error) {
	respMap := initiateTaskMap();
	val, ok := respMap[id];
	if ok {
		return *val, nil
	} else {
		return *val, errhand.TaskNotFound
	}
}

func AcceptPriceRequest(w http.ResponseWriter, r *http.Request) {
	priceRequest := PriceRequest{}
	_ = json.NewDecoder(r.Body).Decode(&priceRequest)
    w.WriteHeader(http.StatusAccepted)
	taskId := TaskCounter.getTaskId();
	go proceed(taskId, priceRequest.Isin, priceRequest.Underlying, priceRequest.Volatility, nil)
    response := AsyncResponse{"price/"+strconv.FormatUint(taskId,10)}
    json.NewEncoder(w).Encode(response);
}

func WaitPriceRequest(w http.ResponseWriter, r *http.Request) {
	priceRequest := PriceRequest{}
	_ = json.NewDecoder(r.Body).Decode(&priceRequest)
	taskId := TaskCounter.getTaskId();
	signalChan := make(chan int)
	go proceed(taskId, priceRequest.Isin, priceRequest.Underlying, priceRequest.Volatility, signalChan)
	if signal := <- signalChan; signal == -1 {
		http.Error(w, errhand.TaskCanselledByTimeOut.Error(), http.StatusServiceUnavailable)
        return
	}
    response, err := getTaskState(taskId)
    if err == errhand.TaskNotFound {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    } else if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
	}
    json.NewEncoder(w).Encode(response);
}

func ReturnTaskRequest(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r);
	var param string = params["id"]
	taskId, err := strconv.ParseUint(param, 10, 64)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
	response, err := getTaskState(taskId)
    if err == errhand.TaskNotFound {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    } else if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
	}
    json.NewEncoder(w).Encode(response);
}

func LogWrapper(h func(http.ResponseWriter, *http.Request)) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    cookie, _ := r.Cookie("clientId")
	if cookie == nil {
		http.Error(w, "Client cookie not found", 400)
		return
	} else {
		allowed := access.AccessRateControl(cookie.Value)
		if !allowed {
			http.Error(w, "Too many requests", 400)
			return
		}
	}
    h(w, r) // call original
  })
}