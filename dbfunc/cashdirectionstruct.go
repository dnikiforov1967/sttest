package dbfunc

//CacheDirection representation
type CashDirectionStruct struct {
	Path string `json:"path"`
	CashType string `json:"type"`
	PaymentStruct `json:"payment"`
}
