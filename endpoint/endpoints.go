package endpoint

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/microbusinesses/ServiceDiscoveryService/business/contract"
	"github.com/microbusinesses/ServiceDiscoveryService/endpoint/message"
	"golang.org/x/net/context"
)

func createResolveServiceEndpoint(service contract.ServiceDiscoveryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(message.ResolveServiceRequest)

		serviceAddresses, err := service.ResolveService(req.ServiceName)

		if err != nil {
			return message.ResolveServiceResponse{nil, err.Error()}, err
		}

		return message.ResolveServiceResponse{serviceAddresses, ""}, nil

	}
}
