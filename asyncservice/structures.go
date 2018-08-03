package asyncservice

import "sync"

//Structure of async response
type AsyncResponse struct {
    ResourcePath string `json:"resource"`
}

//Structure of price request
type PriceRequest struct {
    Isin string `json:"isin"`
	Underlying float64 `json:"underlying"`
	Volatility float64 `json:"volatility"`
}

//Structure of task counter
type taskCounter struct {
	counter uint64
}

//Structure of task response
type TaskResponse struct {
	Id uint64 `json:"id"`
	Isin string `json:"isin"`
    Status string `json:"status"`
	Price float64 `json:"price"`
	PriceDate string `json:"date"`
	taskMutex sync.RWMutex
}


