package transport

import (
	"encoding/json"
	"net/http"

	"golang.org/x/net/context"

	"github.com/microbusinesses/ServiceDiscoveryService/endpoint/message"
)

func DecodeResolveServiceRequest(context context.Context, httpRequest *http.Request) (interface{}, error) {
	var request message.ResolveServiceRequest

	if err := json.NewDecoder(httpRequest.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}
