package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "encoding/json"
    "bytes"

    "github.com/gorilla/mux"
    "github.com/stretchr/testify/assert"
    "github.com/dnikiforov1967/sttest/dbfunc"
    "github.com/dnikiforov1967/sttest/rest"
    "github.com/dnikiforov1967/sttest/config"
)

func Router() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/product", rest.CreateProduct).Methods("POST")
    router.HandleFunc("/product/{id}", rest.GetProduct).Methods("GET")
    router.HandleFunc("/product/{id}", rest.UpdateProduct).Methods("PUT")
    router.HandleFunc("/product/{id}", rest.DeleteProduct).Methods("DELETE")
    return router
}

func TestCreateEndpoint(t *testing.T) {

    //Read standard configuration
    config.ReadFromFile(config.ConfigFileName)

    product := dbfunc.GetTestProduct("ID0")
    jsonProduct, _ := json.Marshal(&product)

    requestPost, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(jsonProduct))
    responsePost := httptest.NewRecorder()
    Router().ServeHTTP(responsePost, requestPost)
    assert.Equal(t, 201, responsePost.Code, "OK response is expected")

    requestDelete, _ := http.NewRequest("DELETE", "/product/"+product.Product_id, nil)
    responseDelete := httptest.NewRecorder()
    Router().ServeHTTP(responseDelete, requestDelete)
    assert.Equal(t, 204, responseDelete.Code, "OK response is expected")

}
