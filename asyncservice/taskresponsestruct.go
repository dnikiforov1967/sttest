package asyncservice

import (
    "time"
    "sync"
)

//Structure of task response
type TaskResponseStruct struct {
    Id uint64 `json:"id"`
    Isin string `json:"isin"`
    Status string `json:"status"`
    Price float64 `json:"price"`
    PriceDate string `json:"date"`
    taskMutex sync.RWMutex
}

//Method updates task (row-specific lock)
func (response *TaskResponseStruct) writeData(price float64) {
	response.taskMutex.Lock()
	defer response.taskMutex.Unlock()
	response.Status = StatusCompleted
	response.Price = Round2(price)
	response.PriceDate = time.Now().Format(time.RFC3339)
}

//Method locks task instance in read mode
func (response *TaskResponseStruct) readLock() {
	response.taskMutex.RLock()
}

//Method unlocks task instance in read mode
func (response *TaskResponseStruct) readUnlock() {
	response.taskMutex.RUnlock()
}