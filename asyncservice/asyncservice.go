package asyncservice

import "net/http"
import "encoding/json"

func AcceptPriceRequest(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusAccepted)
    response := AsyncResponse{"price/1234"}
    json.NewEncoder(w).Encode(response);
}
