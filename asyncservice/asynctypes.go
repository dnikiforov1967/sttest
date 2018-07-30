package asyncservice

import "sync/atomic"

type AsyncResponse struct {
    ResourcePath string `json:"resource"`
}

type PriceRequest struct {
    Isin string `json:"isin"`
	Underlying float64 `json:"underlying"`
	Volatility float64 `json:"volatility"`
}

type taskCounter struct {
	counter uint64
}

type TaskResponse struct {
	Id uint64 `json:"id"`
	Isin string `json:"isin"`
    Status string `json:"status"`
	Price float64 `json:"price"`
	PriceDate string `json:"date"`
}


func (tc *taskCounter) getTaskId() uint64 {
	return atomic.AddUint64(&tc.counter,1)
}

var TaskCounter taskCounter = taskCounter{0}

const StatusCompleted string = "COMPLETED"
const StatusInProgress string = "IN PROGRESS"
const StatusFailed string = "FAILED"
const StatusTimedOut string = "TIMED OUT"