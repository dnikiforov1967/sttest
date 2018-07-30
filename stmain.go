package main

import "log"
import "net/http"
import "github.com/gorilla/mux"
import "./rest"
import "./asyncservice"

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/product", rest.CreateProduct).Methods("POST")
    router.HandleFunc("/product/{id}", rest.GetProduct).Methods("GET")
    router.HandleFunc("/product/{id}", rest.UpdateProduct).Methods("PUT")
    router.HandleFunc("/product/{id}", rest.DeleteProduct).Methods("DELETE")
    router.HandleFunc("/price", asyncservice.AcceptPriceRequest).Methods("POST")
	router.HandleFunc("/price/{id}", asyncservice.ReturnTaskRequest).Methods("GET")
    log.Fatal(http.ListenAndServe(":8080", router))
}
