package main

import "fmt"
import "./dbfunc"

func main() {
	fmt.Printf("ST test started\n");
        product := dbfunc.Product{0, "S","T1","Y",false,"2017-01-01","2017-01-01"}
        product.InsertProduct();
        fmt.Printf("ID IS %d\n",product.Id);
        product.Name="XXX"
        product.UpdateProduct();

        fetched := dbfunc.Product{}
        fetched.Product_id = product.Product_id
        fetched.FetchProductByProductId()
        fmt.Println(fetched.Name);
        fmt.Println(fetched.Quanto);
        product.DeleteProductByProductId();
}
