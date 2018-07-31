package access

import "time"

type AccessLimitStruct struct {
	ClientId string `json:"clientId"`
	Limiit int32 `json:"limit"`
}

type accessTrackingStruct struct {
	incomedRequests int64
	firstIncomeTime time.Time
}

