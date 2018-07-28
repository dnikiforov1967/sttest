package rest

import "../dbfunc"
import "net/http"
import "../errhand"
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
    err := product.FetchProductByProductId()
    if err == errhand.ErrProdNotFound {
        http.Error(w, err.Error(), 404)
        return
    } else if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    json.NewEncoder(w).Encode(product);
}