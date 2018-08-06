package config

import "sync"
import "sync/atomic"

var TimeOut int64
var Database atomic.Value

var configMutex sync.RWMutex