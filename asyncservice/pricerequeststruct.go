package asyncservice

//Structure of price request
type PriceRequestStruct struct {
    Isin string `json:"isin"`
	Underlying float64 `json:"underlying"`
	Volatility float64 `json:"volatility"`
}
