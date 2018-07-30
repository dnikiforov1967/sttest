package asyncservice

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
