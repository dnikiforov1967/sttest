package config

import (
	"io/ioutil"
	"log"
	"encoding/json"
    "github.com/dnikiforov1967/accesslib"
	"sync/atomic"
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
