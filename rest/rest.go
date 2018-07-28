package rest

import "../dbfunc"
import "net/http"
import "encoding/json"
import "github.com/gorilla/mux"

func CreateProduct(w http.ResponseWriter, r *http.Request) {
    var product dbfunc.Product 
    _ = json.NewDecoder(r.Body).Decode(&product)
    product.InsertProduct()
    json.NewEncoder(w).Encode(product);
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r) 
    var product dbfunc.Product
    product.Product_id = params["id"]
    product.FetchProductByProductId()
    json.NewEncoder(w).Encode(product);
}