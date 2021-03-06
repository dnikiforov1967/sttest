package rest

import "github.com/dnikiforov1967/sttest/dbfunc"
import "net/http"
import "github.com/dnikiforov1967/sttest/errhand"
import "encoding/json"
import "github.com/gorilla/mux"

func CreateProduct(w http.ResponseWriter, r *http.Request) {
    var product dbfunc.ProductStruct 
    _ = json.NewDecoder(r.Body).Decode(&product)
    err := product.InsertProduct()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(product);
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r) 
    var product dbfunc.ProductStruct
    product.Product_id = params["id"]
    err := product.FetchProductByProductId()
    if err == errhand.ProdNotFound {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    } else if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(product);
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
    var product dbfunc.ProductStruct
    params := mux.Vars(r)
    var origId string = params["id"]
    _ = json.NewDecoder(r.Body).Decode(&product)
    err := product.UpdateProduct(origId)
    if err == errhand.ProdNotFound {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    } else if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(product);
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
    var product dbfunc.ProductStruct
    params := mux.Vars(r)
    var origId string = params["id"]
    product.Product_id = origId
    err := product.DeleteProductByProductId()
    if err == errhand.ProdNotFound {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    } else if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}