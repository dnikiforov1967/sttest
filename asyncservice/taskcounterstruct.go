package asyncservice

import "sync/atomic"

//Structure of task counter
type taskCounterStruct struct {
	counter uint64
}

//Method returns newly generated task Id
func (tc *taskCounterStruct) getTaskId() uint64 {
	return atomic.AddUint64(&tc.counter,1)
}

