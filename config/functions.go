package config

import (
	"io/ioutil"
	"log"
	"strconv"
	"encoding/json"
    "github.com/dnikiforov1967/accesslib"
	"sync/atomic"
	"net/http"
	"github.com/gorilla/mux"
)

func ReadFromFile(fileName string) {
	conf := &ConfigStruct{}
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = json.Unmarshal(file, conf)
        if err != nil {
            log.Fatal(err)
        }
		atomic.StoreInt64(&TimeOut, conf.Timeout)
		Database = conf.Database
        for _, value := range conf.Limits {
            accesslib.ClientLimits[value.ClientId] = value.Limit
        }
}

func SetTimeout(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r);
	var param string = params["timeout"]
	timeout, err := strconv.ParseInt(param, 10, 64)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
	atomic.StoreInt64(&TimeOut, timeout)
	w.WriteHeader(http.StatusAccepted)
}
