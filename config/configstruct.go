package config

import (
	"io/ioutil"
	"log"
	"encoding/json"
)

type ConfigStruct struct {
	Timeout int64 `json:"timeout"`
}

var GlobalConfig ConfigStruct = ConfigStruct{}

func (conf *ConfigStruct) ReadFromFile(fileName string) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err.Error())
	}
	json.Unmarshal(file, conf)
}