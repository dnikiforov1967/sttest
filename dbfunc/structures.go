package dbfunc

type Product struct {
    id int64
    Name string `json:"name"`
    Product_id string `json:"product_id"`
    Category string `json:"category"`
    Quanto bool `json:"quanto"`
    CreationDate string `json:"creationDate"`
    ExpirationDate string `json:"expirationDate"`
    Terms TermsStruct `json:"terms"`
}

type TermsStruct struct {
    Events []Event `json:"events"`
}

type Event struct {
    id int64
    parent_id int64
    EventType string `json:"type"`
	Terminal bool `json:"terminal"`
	ExecutionStruct `json:"execution"`
	CashDirection `json:"cashDirection"`
}

type ExecutionStruct struct {
	OnStruct `json:"on"`
	Origin string `json:"origin"`
	ExecType string `json:"type"`
}

type OnStruct struct {
	Kind string `json:"kind"`
}

type CashDirection struct {
	Path string `json:"path"`
	CashType string `json:"type"`
	Payment `json:"payment"`
}

type Payment struct {
	PaymentType string `json:"type"`
	Method string `json:"method"`
	AlgorithmId string `json:"algorithmId"`
}