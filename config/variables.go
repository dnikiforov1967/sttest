package config

import "sync"

var TimeOut int64
var Database string

var configMutex sync.RWMutex