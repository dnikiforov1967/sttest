package asyncservice

import "sync/atomic"

func (tc *taskCounter) getTaskId() uint64 {
	return atomic.AddUint64(&tc.counter,1)
}

