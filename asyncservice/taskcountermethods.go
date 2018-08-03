package asyncservice

import "sync/atomic"

//Method returns newly generated task Id
func (tc *taskCounter) getTaskId() uint64 {
	return atomic.AddUint64(&tc.counter,1)
}

