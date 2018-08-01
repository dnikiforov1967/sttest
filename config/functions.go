package config

import (
	"io/ioutil"
	"log"
	"encoding/json"
    "../access"
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
            access.ClientLimits[value.ClientId] = value.Limit
        }
}
