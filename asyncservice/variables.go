package asyncservice

import "sync"

var TaskCounter taskCounter = taskCounter{0}

var ( 
	taskMap map[uint64]*TaskResponse
	mapLock sync.Mutex
)