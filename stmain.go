package main

import "log"
import "net/http"
import "github.com/gorilla/mux"
import "./rest"
import "./asyncservice"
import "./config"
//import "./access"

func main() {
	
	//Read standard configuration
	config.ReadFromFile("./config.json")

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/product", rest.CreateProduct).Methods("POST")
    router.HandleFunc("/product/{id}", rest.GetProduct).Methods("GET")
    router.HandleFunc("/product/{id}", rest.UpdateProduct).Methods("PUT")
    router.HandleFunc("/product/{id}", rest.DeleteProduct).Methods("DELETE")

	priceRtr := router.PathPrefix("/price").Subrouter().StrictSlash(true)
    priceRtr.Handle("", asyncservice.LogWrapper(asyncservice.AcceptPriceRequest)).Methods("POST")
	priceRtr.Handle("/{id}", asyncservice.LogWrapper(asyncservice.ReturnTaskRequest)).Methods("GET")
	priceRtr.Handle("/synch/wait", asyncservice.LogWrapper(asyncservice.WaitPriceRequest)).Methods("POST")

    log.Fatal(http.ListenAndServe(":8080", router))
}
