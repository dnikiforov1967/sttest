package asyncservice

import "net/http"
import "encoding/json"
import "strconv"
import "time"
import "math"
import "sync"
import "sync/atomic"
import "github.com/gorilla/mux"
import "github.com/dnikiforov1967/sttest/errhand"
import "github.com/dnikiforov1967/sttest/config"
import "github.com/dnikiforov1967/accesslib"

//function safely instantiate new task map structure
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

//function rounds float64 to 2 dec. points (bug in go_wrapper does not allow math.Round() to use)
func Round2(x float64) float64 {
	x = x *100
    t := math.Trunc(x)
    if math.Abs(x-t) >= 0.5 {
        t = t + math.Copysign(1, x)
	}
    return t/100
}

//Function executes task. Taks execution takes about 5 sec
func proceed(id uint64, isin string, underlying float64, volatility float64, signalChan chan int) {
	respMap := initiateTaskMap();
	response := TaskResponse{id, isin, StatusInProgress, 0, "", sync.RWMutex{}}
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
	response.writeData(underlying*volatility*1000)
	if signalChan != nil {
		signalChan <- 0
	}
}

//Function checks if timeout expired
func checkTimeOut(initTime *time.Time) bool {
	duration := time.Since(*initTime)
	var millisec int64 = duration.Nanoseconds()/1000000
	timeOut := atomic.LoadInt64(&config.TimeOut)
	if millisec >= timeOut {
		return true
	}
	return false
}

//Function returns task state
func getTaskState(id uint64) (TaskResponse, error) {
	respMap := initiateTaskMap();
	val, ok := respMap[id];
	val.readLock()
	defer val.readUnlock()
	if ok {
		return *val, nil
	} else {
		return *val, errhand.TaskNotFound
	}
}

//Function proceeds price request. It returns path to task resource for later check
func AcceptPriceRequest(w http.ResponseWriter, r *http.Request) {
	priceRequest := PriceRequest{}
	_ = json.NewDecoder(r.Body).Decode(&priceRequest)
    w.WriteHeader(http.StatusAccepted)
	taskId := TaskCounter.getTaskId();
	go proceed(taskId, priceRequest.Isin, priceRequest.Underlying, priceRequest.Volatility, nil)
    response := AsyncResponse{"price/"+strconv.FormatUint(taskId,10)}
    json.NewEncoder(w).Encode(response);
}

//Function proceeds price request in sync.  
func WaitPriceRequest(w http.ResponseWriter, r *http.Request) {
	priceRequest := PriceRequest{}
	_ = json.NewDecoder(r.Body).Decode(&priceRequest)
	taskId := TaskCounter.getTaskId();
	signalChan := make(chan int)
	go proceed(taskId, priceRequest.Isin, priceRequest.Underlying, priceRequest.Volatility, signalChan)
	if signal := <- signalChan; signal == -1 {
		http.Error(w, errhand.TaskCancelledByTimeOut.Error(), http.StatusServiceUnavailable)
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

//Function returns the state of task
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

//Function checks the rate limit before call of requests
func LogWrapper(h func(http.ResponseWriter, *http.Request)) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    cookie, _ := r.Cookie("clientId")
	if cookie == nil {
		http.Error(w, "Client cookie not found", 400)
		return
	} else {
		allowed := accesslib.AccessRateControl(cookie.Value)
		if !allowed {
			http.Error(w, "Too many requests", 400)
			return
		}
	}
    h(w, r) // call original
  })
}