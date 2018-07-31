package access

var rateLimitMap map[string]*accessTrackingStruct = make(map[string]*accessTrackingStruct)
