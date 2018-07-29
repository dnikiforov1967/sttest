package asyncservice

import "net/http"

func AcceptPriceRequest(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusAccepted)
}
