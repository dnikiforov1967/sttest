package asyncservice

import (
	"sync"
)

//This variable is responsible to keep information about price calculation tasks
var taskIdGenerator taskCounterStruct = taskCounterStruct{0}

//map-specific lock entities
var taskMap taskMapStruct = taskMapStruct{make(map[uint64]*TaskResponseStruct), sync.RWMutex{}}
