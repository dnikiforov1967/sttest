package main

import "log"
import "net/http"
import "github.com/gorilla/mux"
import "github.com/dnikiforov1967/sttest/rest"
import "github.com/dnikiforov1967/sttest/asyncservice"
import "github.com/dnikiforov1967/sttest/config"

//Standard Golang entry point
//Application is designed to proceed web requests from outer clients
func main() {
	
    //Read standard configuration
    config.ReadFromFile(config.ConfigFileName)

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/product", rest.CreateProduct).Methods("POST")
    router.HandleFunc("/product/{id}", rest.GetProduct).Methods("GET")
    router.HandleFunc("/product/{id}", rest.UpdateProduct).Methods("PUT")
    router.HandleFunc("/product/{id}", rest.DeleteProduct).Methods("DELETE")

	priceRtr := router.PathPrefix("/price").Subrouter().StrictSlash(true)
    priceRtr.Handle("", asyncservice.LogWrapper(asyncservice.AcceptPriceRequest)).Methods("POST")
	priceRtr.Handle("/{id}", asyncservice.LogWrapper(asyncservice.ReturnTaskRequest)).Methods("GET")
	priceRtr.Handle("/synch/wait", asyncservice.LogWrapper(asyncservice.WaitPriceRequest)).Methods("POST")

	configRtr := router.PathPrefix("/config").Subrouter().StrictSlash(true)
	configRtr.HandleFunc("/timeout/{timeout}", config.SetTimeout).Methods("PATCH")
	configRtr.HandleFunc("/rate/{clientId}/{rateLimit}", config.SetRateLimit).Methods("PATCH")
	configRtr.HandleFunc("", config.GetConfiguration).Methods("GET")

    log.Fatal(http.ListenAndServe(":8080", router))
}
