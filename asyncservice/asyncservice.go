package asyncservice

import "net/http"
import "encoding/json"
import "strconv"

func AcceptPriceRequest(w http.ResponseWriter, r *http.Request) {
	priceRequest := PriceRequest{}
	_ = json.NewDecoder(r.Body).Decode(&priceRequest)
    w.WriteHeader(http.StatusAccepted)
	taskId := TaskCounter.getTaskId();
    response := AsyncResponse{"price/"+strconv.FormatUint(taskId,10)}
    json.NewEncoder(w).Encode(response);
}
