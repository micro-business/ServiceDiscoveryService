package transport

import (
	"encoding/json"
	"net/http"
)

func EncodeResolveServiceResponse(w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

	return json.NewEncoder(w).Encode(response)
}
