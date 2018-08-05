package asyncservice

import (
	"sync"
	"sync/atomic"
)

//This variable is responsible to keep information about price calculation tasks
var taskIdGenerator taskCounterStruct = taskCounterStruct{0}

//map-specific lock entities
var (
	mapAccess atomic.Value
	mapLock sync.RWMutex
)
