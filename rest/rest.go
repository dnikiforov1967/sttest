package rest

import "../dbfunc"
import "net/http"
import "encoding/json"

func CreateProduct(w http.ResponseWriter, r *http.Request) {
    var product dbfunc.Product 
    _ = json.NewDecoder(r.Body).Decode(&product)
    json.NewEncoder(w).Encode(product);
}
