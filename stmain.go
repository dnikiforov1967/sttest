package main

import "log"
import "net/http"
import "github.com/gorilla/mux"
import "./rest"
import "./asyncservice"
import "./config"

func main() {
	
	//Read standard configuration
	config.GlobalConfig.ReadFromFile("./config.json")

    router := mux.NewRouter().StrictSlash(true)
	priceRtr := router.PathPrefix("/price").Subrouter().StrictSlash(true)
    router.HandleFunc("/product", rest.CreateProduct).Methods("POST")
    router.HandleFunc("/product/{id}", rest.GetProduct).Methods("GET")
    router.HandleFunc("/product/{id}", rest.UpdateProduct).Methods("PUT")
    router.HandleFunc("/product/{id}", rest.DeleteProduct).Methods("DELETE")
    priceRtr.HandleFunc("", asyncservice.AcceptPriceRequest).Methods("POST")
	priceRtr.HandleFunc("/{id}", asyncservice.ReturnTaskRequest).Methods("GET")
	priceRtr.HandleFunc("/synch/wait", asyncservice.WaitPriceRequest).Methods("POST")
    log.Fatal(http.ListenAndServe(":8080", router))
}
