// Defines the service discovery contract.
package contract

import "github.com/microbusinesses/Micro-Businesses-Core/system"

// The service discovery service contract, it can resolve addresses for the available services in the network.
type ServiceDiscoveryService interface {
	// Resovle the provided service name by returning the list of services providing the same functionality.
	// tenantId: Mandatory. The unique identifier of the tenant owning the address.
	// applicationId: Mandatory. The unique identifier of the tenant's application will be owning the address.
	// serviceName: Mandatory. The name of the service to resolve
	// Returns either the collection if available service addresses or error if something goes wrong.
	ResolveService(tenantId, applicationId system.UUID, serviceName string) ([]string, error)
}
