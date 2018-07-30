package asyncservice

import "sync"

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
	taskMap map[uint64]TaskResponse
	mapLock sync.Mutex
)

func initiateTaskMap() map[uint64]TaskResponse {
	if (taskMap != nil) {
		return taskMap
	} else {
		    mapLock.Lock()
			defer mapLock.Unlock()
			if (taskMap != nil) {
				return taskMap
			} else {
				taskMap = make(map[uint64]TaskResponse)
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

func proceed(id uint64, isin string) {
	respMap := initiateTaskMap();
	response := TaskResponse{id, isin, StatusInProgress, 0, ""}
	respMap[id] = response;
}

func getTaskState(id uint64) (TaskResponse, error) {
	respMap := initiateTaskMap();
	val, ok := respMap[id];
	if ok {
		return val, nil
	} else {
		return val, TaskNotFound
	}
}