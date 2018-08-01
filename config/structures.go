package config

import "../access"

type ConfigStruct struct {
	Timeout int64 `json:"timeout"`
    Limits []access.AccessLimitStruct `json:"limits"`
}
