package dbfunc

func GetTestProduct(pid string) Product {
    payment := Payment{"a","b","c"}
    on := On{"X"}
    execution := Execution{on,"Y","Type"}
    cashDirection := CashDirection{"Pay","Cash",payment}
    event := Event{0,0,"EXECUTION",true,execution,cashDirection}
    termsStruct := TermsStruct{}
    termsStruct.Events = append(termsStruct.Events, event)
    product := Product{0,"Name",pid,"Cat",true,"","2018-01-01",termsStruct}
    return product
}