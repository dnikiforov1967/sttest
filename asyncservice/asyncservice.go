package asyncservice

import "net/http"
import "encoding/json"
import "strconv"
import "github.com/gorilla/mux"

func AcceptPriceRequest(w http.ResponseWriter, r *http.Request) {
	priceRequest := PriceRequest{}
	_ = json.NewDecoder(r.Body).Decode(&priceRequest)
    w.WriteHeader(http.StatusAccepted)
	taskId := TaskCounter.getTaskId();
	go proceed(taskId, priceRequest.Isin, nil)
    response := AsyncResponse{"price/"+strconv.FormatUint(taskId,10)}
    json.NewEncoder(w).Encode(response);
}

func WaitPriceRequest(w http.ResponseWriter, r *http.Request) {
	priceRequest := PriceRequest{}
	_ = json.NewDecoder(r.Body).Decode(&priceRequest)
	taskId := TaskCounter.getTaskId();
	signalChan := make(chan int)
	go proceed(taskId, priceRequest.Isin, signalChan)
	if signal := <- signalChan; signal == -1 {
		http.Error(w, TaskCanselledByTimeOut.Error(), http.StatusServiceUnavailable)
        return
	}
    response, err := getTaskState(taskId)
    if err == TaskNotFound {
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
    if err == TaskNotFound {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    } else if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
	}
    json.NewEncoder(w).Encode(response);
}