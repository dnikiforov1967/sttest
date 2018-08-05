package asyncservice

//Structure of price request
type PriceRequest struct {
    Isin string `json:"isin"`
	Underlying float64 `json:"underlying"`
	Volatility float64 `json:"volatility"`
}
