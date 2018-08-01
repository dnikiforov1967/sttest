package config

import "../access"

type configStruct struct {
	timeout int64 `json:"timeout"`
    limits []access.AccessLimitStruct `json:"limits"`
}
