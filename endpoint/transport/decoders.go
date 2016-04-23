package transport

import (
	"encoding/json"
	"net/http"

	"github.com/microbusinesses/ServiceDiscoveryService/endpoint/message"
)

func DecodeResolveServiceRequest(httpRequest *http.Request) (interface{}, error) {
	var request message.ResolveServiceRequest

	if err := json.NewDecoder(httpRequest.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}
