package config

type ConfigStruct struct {
	timeout int `json:"timeout"`
}

func (conf *ConfigStruct) readFromFile(fileName string) error {
	return nil
}