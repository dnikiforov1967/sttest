package dbfunc

//Product top representation
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

//Terms representation
type TermsStruct struct {
    Events []Event `json:"events"`
}

//Event representation
type Event struct {
    id int64
    parent_id int64
    EventType string `json:"type"`
	Terminal bool `json:"terminal"`
	Execution `json:"execution"`
	CashDirection `json:"cashDirection"`
}

//Execution representation
type Execution struct {
	On `json:"on"`
	Origin string `json:"origin"`
	ExecType string `json:"type"`
}

//On representation
type On struct {
	Kind string `json:"kind"`
}

//CacheDirection representation
type CashDirection struct {
	Path string `json:"path"`
	CashType string `json:"type"`
	Payment `json:"payment"`
}

//Payment representation
type Payment struct {
	PaymentType string `json:"type"`
	Method string `json:"method"`
	AlgorithmId string `json:"algorithmId"`
}