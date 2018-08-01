package config

import "github.com/dnikiforov1967/accesslib"

type ConfigStruct struct {
	Timeout int64 `json:"timeout"`
    Limits []accesslib.AccessLimitStruct `json:"limits"`
}
