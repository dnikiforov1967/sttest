package config

import (
	"io/ioutil"
	"log"
	"encoding/json"
    "../access"
)

func ReadFromFile(fileName string) {
	conf := &configStruct{}
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = json.Unmarshal(file, conf)
        if err != nil {
            log.Fatal(err)
        }
		TimeOut = conf.timeout
        for _, value := range conf.limits {
            access.ClientLimits[value.ClientId] = value.Limit
        }
}
