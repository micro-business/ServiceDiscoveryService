package endpoint

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/microbusinesses/Micro-Businesses-Core/system"
	"github.com/microbusinesses/ServiceDiscovery/business/contract"
	"github.com/microbusinesses/ServiceDiscovery/endpoint/message"
	"golang.org/x/net/context"
)

func createResolveServiceEndpoint(service contract.ServiceDiscoveryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(message.ResolveServiceRequest)
		tenantId, _ := system.ParseUUID("02365c33-43d5-4bf8-b220-25563443960b")
		applicationId, _ := system.ParseUUID("02365c33-43d5-4bf8-b220-25563443960c")

		serviceAddresses, err := service.ResolveService(tenantId, applicationId, req.ServiceName)

		if err != nil {
			return message.ResolveServiceResponse{nil, err.Error()}, err
		} else {
			return message.ResolveServiceResponse{serviceAddresses, ""}, nil
		}
	}
}
