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

func (conf *ConfigStruct) ReadFromFile(fileName string) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = json.Unmarshal(file, conf)
        if err != nil {
            log.Fatal(err)
        }
        for _, value := range conf.Limits {
            access.ClientLimits[value.ClientId] = value.Limit
        }
}