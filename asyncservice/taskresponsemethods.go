package asyncservice

import "time"

//Method updates task (row-specific lock)
func (response *TaskResponse) writeData(price float64) {
	response.taskMutex.Lock()
	defer response.taskMutex.Unlock()
	response.Status = StatusCompleted
	response.Price = Round2(price)
	response.PriceDate = time.Now().Format(time.RFC3339)
}

func (response *TaskResponse) readLock() {
	response.taskMutex.RLock()
}

func (response *TaskResponse) readUnlock() {
	response.taskMutex.RUnlock()
}