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
            accesslib.WriteLimit(value.ClientId,value.Limit)
        }
}

func WriteToFile(fileName string, conf *ConfigStruct) {
	bytes, err := json.Marshal(conf);
	if err != nil {
		log.Fatal(err.Error())
	}
	err = ioutil.WriteFile(fileName, bytes, 0644)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func updateConfig() {
	configMutex.Lock()
	defer configMutex.Unlock()
	conf := makeJsonConfig()
	WriteToFile(ConfigFileName, &conf)
}

func SetTimeout(w http.ResponseWriter, r *http.Request) {
	defer updateConfig()
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

func SetRateLimit(w http.ResponseWriter, r *http.Request) {
	defer updateConfig()
	params := mux.Vars(r);
	clientId := params["clientId"]
	var param string = params["rateLimit"]
	rateLimit, err := strconv.ParseInt(param, 10, 64)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
	accesslib.WriteLimit(clientId, rateLimit)
	w.WriteHeader(http.StatusAccepted)
}

func makeJsonConfig() ConfigStruct {
	conf := ConfigStruct{}
	conf.Database = Database
	conf.Timeout = TimeOut
	conf.Limits = []accesslib.AccessLimitStruct{}
	for key, value := range accesslib.ReadLimits() {
		conf.Limits = append(conf.Limits, accesslib.AccessLimitStruct{key, value})
	}
	return conf
}

func GetConfiguration(w http.ResponseWriter, r *http.Request) {
	conf := makeJsonConfig()
	json.NewEncoder(w).Encode(&conf);
}