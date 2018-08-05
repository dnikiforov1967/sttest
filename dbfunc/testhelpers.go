package dbfunc

import "testing"

func GetTestProduct(pid string) ProductStruct {
    payment := PaymentStruct{"a","b","c"}
    on := OnStruct{"X"}
    execution := ExecutionStruct{on,"Y","Type"}
    cashDirection := CashDirectionStruct{"Pay","Cash",payment}
    event := EventStruct{0,0,"EXECUTION",true,execution,cashDirection}
    termsStruct := TermsStruct{}
    termsStruct.Events = append(termsStruct.Events, event)
    product := ProductStruct{0,"Name",pid,"Cat",true,"","2018-01-01",termsStruct}
    return product
}

func CleanupDatabaseBeforeTest(t *testing.T) {
    db, err := openLocalDb()
    if err != nil {
        t.Errorf(err.Error())
        return
    }
    defer db.Close()
    tx, err := openTrans(db)
    if err != nil {
        t.Errorf(err.Error())
        return
    }
    defer tx.Commit()
    _, err = tx.Exec("delete from events")
    if err != nil {
        t.Errorf(err.Error())
        return
    }
    _, err = tx.Exec("delete from products")
    if err != nil {
        t.Errorf(err.Error())
        return
    }

}