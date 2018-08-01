package config

import (
	"io/ioutil"
	"log"
	"encoding/json"
    "github.com/dnikiforov1967/accesslib"
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
		TimeOut = conf.Timeout
        for _, value := range conf.Limits {
            accesslib.ClientLimits[value.ClientId] = value.Limit
        }
}
