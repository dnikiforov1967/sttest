package asyncservice

import "sync"
import "time"
import "fmt"
import "math"

type TaskResponse struct {
	Id uint64 `json:"id"`
	Isin string `json:"isin"`
    Status string `json:"status"`
	Price float64 `json:"price"`
	PriceDate string `json:"date"`
}

const StatusCompleted string = "COMPLETED"
const StatusInProgress string = "IN PROGRESS"
const StatusFailed string = "FAILED"
const StatusTimedOut string = "TIMED OUT"

var ( 
	taskMap map[uint64]*TaskResponse
	mapLock sync.Mutex
)

func initiateTaskMap() map[uint64]*TaskResponse {
	if (taskMap != nil) {
		return taskMap
	} else {
		    mapLock.Lock()
			defer mapLock.Unlock()
			if (taskMap != nil) {
				return taskMap
			} else {
				taskMap = make(map[uint64]*TaskResponse)
				return taskMap
			}
	}
}

type asyncError struct {
	message string
}

func (err asyncError) Error() string {
	return err.message
}

var TaskNotFound asyncError = asyncError{"Task not found"}
var TaskCanselledByTimeOut asyncError = asyncError{"Task cancelled by timeout"}

func proceed(id uint64, isin string, underlying float64, volatility float64, signalChan chan int) {
	respMap := initiateTaskMap();
	response := TaskResponse{id, isin, StatusInProgress, 0, ""}
	respMap[id] = &response;

	initTime := time.Now()
	
	//Summary time ~ 5 sec
	for i := 0; i<10; i++ {
		//Here we simulate steps of price calculation
		fmt.Println("Step ", i)
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
	var millisec int64 = duration.Nanoseconds()
	var limit int64 = 7000000000;
	fmt.Println("Milli ", millisec)
	if millisec >= limit {
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
		return *val, TaskNotFound
	}
}