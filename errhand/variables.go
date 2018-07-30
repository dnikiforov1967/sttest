package errhand

import "errors"

var ErrProdNotFound error = errors.New("Product not found")