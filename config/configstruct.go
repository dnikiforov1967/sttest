package config

import (
	"io/ioutil"
	"log"
	"encoding/json"
        "../access"
)

type ConfigStruct struct {
	Timeout int64 `json:"timeout"`
        Limits []access.AccessLimitStruct `json:"limits"`
}

var GlobalConfig ConfigStruct = ConfigStruct{}
var ClientLimits map[string]int64 = make(map[string]int64)

func (conf *ConfigStruct) ReadFromFile(fileName string) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err.Error())
	}
	json.Unmarshal(file, conf)
        for _, value := range conf.Limits {
            ClientLimits[value.ClientId] = value.Limit
        }
}