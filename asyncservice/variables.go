package asyncservice

import "sync"
import "sync/atomic"

var TaskCounter taskCounter = taskCounter{0}

var (
	mapAccess atomic.Value
	mapLock sync.Mutex
)
