package main

import "log"
import "net/http"
import "github.com/gorilla/mux"
import "./rest"

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/product", rest.CreateProduct).Methods("POST")
    router.HandleFunc("/product/{id}", rest.GetProduct).Methods("GET")
    router.HandleFunc("/product/{id}", rest.UpdateProduct).Methods("PUT")
    log.Fatal(http.ListenAndServe(":8080", router))
}
