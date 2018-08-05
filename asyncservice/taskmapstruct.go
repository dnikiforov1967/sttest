package asyncservice

import "sync"

type taskMapStruct struct {
    internalMap map[uint64]*TaskResponseStruct
    mapLock sync.RWMutex    
}

func (obj *taskMapStruct) writeToMap(id uint64, response *TaskResponseStruct) {
    obj.mapLock.Lock()
    defer obj.mapLock.Unlock()
    obj.internalMap[id] = response
}

func (obj *taskMapStruct) readFromMap(id uint64) (*TaskResponseStruct, bool) {
    obj.mapLock.RLock()
    defer obj.mapLock.RUnlock()
    val, ok := obj.internalMap[id]
    return val, ok
}