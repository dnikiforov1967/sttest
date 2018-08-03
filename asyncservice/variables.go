package asyncservice

import (
	"sync"
	"sync/atomic"
)

//This variable is responsible to keep information about price calculation tasks
var TaskCounter taskCounter = taskCounter{0}

var (
	mapAccess atomic.Value
	mapLock sync.Mutex
)
