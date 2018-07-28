package main

import "fmt"
import "./dbfunc"

func main() {
	fmt.Printf("ST test started\n");
        product := dbfunc.Product{0, "S","T1","Y",true,"2017-01-01","2017-01-01"}
        product.InsertProduct();
        fmt.Printf("ID IS %d\n",product.Id);
        product.DeleteProductByProductId();
}
