package asyncservice

import "sync/atomic"

//Structure of task counter
type taskCounter struct {
	counter uint64
}

//Method returns newly generated task Id
func (tc *taskCounter) getTaskId() uint64 {
	return atomic.AddUint64(&tc.counter,1)
}

