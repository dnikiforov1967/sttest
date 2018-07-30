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


func (tc *taskCounter) getTaskId() uint64 {
	return atomic.AddUint64(&tc.counter,1)
}

var TaskCounter taskCounter = taskCounter{0}
