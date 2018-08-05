package config

import "github.com/dnikiforov1967/accesslib"

type ConfigStruct struct {
	Database string `json:"database"`
	Timeout int64 `json:"timeout"`
    Limits []accesslib.AccessLimitStruct `json:"limits"`
}
