package errhand

import "errors"

var ProdNotFound error = errors.New("Product not found")
var TaskNotFound error = errors.New("Task not found")
var TaskCanselledByTimeOut error = errors.New("Task cancelled by timeout")