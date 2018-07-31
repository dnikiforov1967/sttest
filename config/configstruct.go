package config

import (
	"io/ioutil"
	"log"
	"encoding/json"
        "../access"
        "fmt"
)

type ConfigStruct struct {
	Timeout int64 `json:"timeout"`
        Limits []access.AccessLimitStruct `json:"limits"`
}

var GlobalConfig ConfigStruct = ConfigStruct{}

func (conf *ConfigStruct) ReadFromFile(fileName string) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err.Error())
	}
	json.Unmarshal(file, conf)
        for _, value := range conf.Limits {
            access.ClientLimits[value.ClientId] = value.Limit
            fmt.Printf("Limit got %d", value.Limit)
        }
}