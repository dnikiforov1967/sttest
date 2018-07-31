package access

import "sync"
import "time"
import "fmt"

var lock sync.Mutex

func AccessRateControl(clientId string) bool {
	lock.Lock();
	defer lock.Unlock()
	returnValue := false;
	val, ok := rateLimitMap[clientId]	
	if !ok {
		fmt.Printf("We initiate rate limit for %s\n",clientId);
		val = &accessTrackingStruct{1, time.Now()}
		rateLimitMap[clientId] = val
		returnValue = true
	} else {
		currTime := time.Now()
		dur := currTime.Sub(val.firstIncomeTime)
		if (dur.Nanoseconds() > 1000000000) {
			val.firstIncomeTime = currTime
			val.incomedRequests = 1
			returnValue = true
		} else {
			val.incomedRequests++
			returnValue = false
		}
	}
	fmt.Printf("Now time is %s, request %d\n", val.firstIncomeTime, val.incomedRequests);
	return returnValue;
}
