package access

import "sync"
import "time"
import "fmt"

var lock sync.Mutex

//Access rate controller should be defended by mutex (at least if we want to
//implement lazy initialization
func AccessRateControl(clientId string, limit int64) bool {
	lock.Lock();
	defer lock.Unlock()
	returnValue := true;
	val, ok := rateLimitMap[clientId]	
	if !ok {
		fmt.Printf("We initiate rate limit for %s\n",clientId);
		val = &accessTrackingStruct{1, time.Now()}
		rateLimitMap[clientId] = val
	} else {
		currTime := time.Now()
		dur := currTime.Sub(val.firstIncomeTime)
		if (dur.Nanoseconds()/1000000 > 1000) {
			val.firstIncomeTime = currTime
			val.incomedRequests = 1
		} else {
			val.incomedRequests++
			//Here should be limit check
			returnValue = !(limit < val.incomedRequests)
		}
	}
	fmt.Printf("Now time is %s, request %d\n", val.firstIncomeTime, val.incomedRequests);
	return returnValue;
}
