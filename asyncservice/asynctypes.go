package asyncservice

type AsyncResponse struct {
    ResourcePath string `json:"resource"`
}

type PriceRequest struct {
    Isin string `json:"isin"`
	Underlying float64 `json:"underlying"`
	Volatility float64 `json:"volatility"`
}