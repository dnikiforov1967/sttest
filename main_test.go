package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "encoding/json"
    "bytes"
    "reflect"

    "github.com/gorilla/mux"
    "github.com/stretchr/testify/assert"
    "github.com/dnikiforov1967/sttest/dbfunc"
    "github.com/dnikiforov1967/sttest/rest"
    "github.com/dnikiforov1967/sttest/config"
    "github.com/dnikiforov1967/sttest/asyncservice"
)

func Router() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/product", rest.CreateProduct).Methods("POST")
    router.HandleFunc("/product/{id}", rest.GetProduct).Methods("GET")
    router.HandleFunc("/product/{id}", rest.UpdateProduct).Methods("PUT")
    router.HandleFunc("/product/{id}", rest.DeleteProduct).Methods("DELETE")
    router.Handle("/price", asyncservice.LogWrapper(asyncservice.AcceptPriceRequest)).Methods("POST")
    return router
}

func PriceRouter(router *mux.Router) *mux.Router {
    priceRtr := router.PathPrefix("/price").Subrouter().StrictSlash(true)
    return priceRtr
}

func TestCreateEndpoint(t *testing.T) {

    //Read standard configuration
    config.ReadFromFile(config.ConfigFileName)

    product := dbfunc.GetTestProduct("ID0")
    jsonProduct, _ := json.Marshal(&product)

    router := Router()

    requestPost, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(jsonProduct))
    responsePost := httptest.NewRecorder()
    router.ServeHTTP(responsePost, requestPost)
    assert.Equal(t, 201, responsePost.Code, "OK response is expected")
    insertedProduct := dbfunc.Product{}
    json.Unmarshal(responsePost.Body.Bytes(), &insertedProduct)

    requestGet, _ := http.NewRequest("GET", "/product/"+product.Product_id, nil)
    responseGet := httptest.NewRecorder()
    router.ServeHTTP(responseGet, requestGet)
    assert.Equal(t, 200, responseGet.Code, "OK response is expected")
    fetchedProduct := dbfunc.Product{}
    json.Unmarshal(responseGet.Body.Bytes(), &fetchedProduct)

    ok := reflect.DeepEqual(fetchedProduct, insertedProduct)
    assert.Equal(t, true, ok, "Products identical")

    insertedProduct.Name = "New name"
    jsonProduct, _ = json.Marshal(&insertedProduct)
    requestPut, _ := http.NewRequest("PUT", "/product/"+product.Product_id, bytes.NewBuffer(jsonProduct))
    responsePut := httptest.NewRecorder()
    router.ServeHTTP(responsePut, requestPut)
    assert.Equal(t, 200, responsePut.Code, "OK response is expected")
    updatedProduct := dbfunc.Product{}
    json.Unmarshal(responsePut.Body.Bytes(), &updatedProduct)

    requestGet, _ = http.NewRequest("GET", "/product/"+product.Product_id, nil)
    responseGet = httptest.NewRecorder()
    router.ServeHTTP(responseGet, requestGet)
    assert.Equal(t, 200, responseGet.Code, "OK response is expected")
    fetchedProduct = dbfunc.Product{}
    json.Unmarshal(responseGet.Body.Bytes(), &fetchedProduct)

    ok = reflect.DeepEqual(fetchedProduct, updatedProduct)
    assert.Equal(t, true, ok, "Products identical")
 
    requestDelete, _ := http.NewRequest("DELETE", "/product/"+product.Product_id, nil)
    responseDelete := httptest.NewRecorder()
    router.ServeHTTP(responseDelete, requestDelete)
    assert.Equal(t, 204, responseDelete.Code, "OK response is expected")

    requestGet, _ = http.NewRequest("GET", "/product/"+product.Product_id, nil)
    responseGet = httptest.NewRecorder()
    router.ServeHTTP(responseGet, requestGet)
    assert.Equal(t, 404, responseGet.Code, "OK response is expected")

    priceRequest := asyncservice.PriceRequest{"ISIN0",0.4, 0.007}
    jsonPriceRequest, _ := json.Marshal(&priceRequest)
    request, _ := http.NewRequest("POST", "/price", bytes.NewBuffer(jsonPriceRequest))
    response := httptest.NewRecorder()
    router.ServeHTTP(response, request)
    assert.Equal(t, 400, response.Code, "OK response is expected")

}
