package errhand

type errCustom struct {
    message string
}

func (e errCustom) Error() string {
    return e.message
}

var ErrProdNotFound errCustom = errCustom{"Product not found"}