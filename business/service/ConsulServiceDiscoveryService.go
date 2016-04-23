package service

import (
	"github.com/microbusinesses/Micro-Businesses-Core/common/diagnostics"
	"github.com/microbusinesses/Micro-Businesses-Core/system"
)

// The consul service discovery uses Consul to resolve available  services in the network.
type ConsulServiceDiscoveryService struct {
}

// Resovle the provided service name by returning the list of services providing the same functionality.
// tenantId: Mandatory. The unique identifier of the tenant owning the address.
// applicationId: Mandatory. The unique identifier of the tenant's application will be owning the address.
// serviceName: Mandatory. The name of the service to resolve
// Returns either the collection if available service addresses or error if something goes wrong.
func (ConsulServiceDiscoveryService ConsulServiceDiscoveryService) ResolveService(tenantId, applicationId system.UUID, serviceName string) ([]string, error) {
	diagnostics.IsNotNilOrEmpty(tenantId, "tenantId", "tenantId must be provided.")
	diagnostics.IsNotNilOrEmpty(applicationId, "applicationId", "applicationId must be provided.")
	diagnostics.IsNotNilOrEmptyOrWhitespace(serviceName, "serviceName", "serviceName must be provided. Cannot be nil, empty or contains whitespace character only.")

	return nil, nil
}
