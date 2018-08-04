package dbfunc

import (
    "testing"
    "reflect"
    "github.com/dnikiforov1967/sttest/config"
    "github.com/dnikiforov1967/sttest/errhand"
)

func TestDb(t *testing.T) {
    config.Database = "../sttest.sqlt"

    payment := Payment{"a","b","c"}
    on := On{"X"}
    execution := Execution{on,"Y","Type"}
    cashDirection := CashDirection{"Pay","Cash",payment}
    event := Event{0,0,"EXECUTION",true,execution,cashDirection}
    termsStruct := TermsStruct{}
    termsStruct.Events = append(termsStruct.Events, event)
    product := Product{0,"Name","ID1","Cat",true,"","2018-01-01",termsStruct}

    err := product.InsertProduct()
    if err != nil {
        t.Errorf("Failed Insert %s\n", err.Error())
        return
    }

    productToCheck := Product{}
    productToCheck.Product_id = product.Product_id

    productToCheck.FetchProductByProductId()
    //TODO - should compare the complete object
    if !reflect.DeepEqual(product,productToCheck) {
        t.Errorf("Incomplete insert %s\n", err.Error())
        return
    }

    product.Name = "New name"
    product.UpdateProduct("ID1")

    productToCheck = Product{}
    productToCheck.Product_id = product.Product_id

    productToCheck.FetchProductByProductId()
    //TODO - should compare the complete object
    if !reflect.DeepEqual(product,productToCheck) {
        t.Errorf("Incomplete insert %s\n", err.Error())
        return
    }    

    err = product.DeleteProductByProductId()
    if err != nil {
        t.Errorf("Failed deleting %s\n", err.Error())
        return
    }
    err = product.FetchProductByProductId()
    if err != errhand.ProdNotFound {
        t.Errorf("Failed deletion")
        return
    }

    db, err := openLocalDb()
    if err != nil {
        t.Errorf("Failed db open")
        return
    }
    defer db.Close()
    row := db.QueryRow("select count(*) from events where parent_id = $1", product.id);
    if err != nil {
        t.Errorf("Events fetch %s",err.Error())
        return
    }
    var cnt int64
    err = row.Scan(&cnt)
    if cnt != 0 {
        t.Errorf("Orphan events")
        return
    }
    
}