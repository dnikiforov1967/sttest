package dbfunc

//Execution representation
type ExecutionStruct struct {
	OnStruct `json:"on"`
	Origin string `json:"origin"`
	ExecType string `json:"type"`
}
