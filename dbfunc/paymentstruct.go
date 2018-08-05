package dbfunc

//Payment representation
type PaymentStruct struct {
	PaymentType string `json:"type"`
	Method string `json:"method"`
	AlgorithmId string `json:"algorithmId"`
}