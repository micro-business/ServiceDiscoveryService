// Package contract defines the service discovery contract.
package contract

// ServiceDiscoveryService contract defines the service that can resolve addresses for the available services in the network.
type ServiceDiscoveryService interface {
	// ResolveService resolves the provided service name by returning the list of services providing the same functionality.
	// serviceName: Mandatory. The name of the service to resolve
	// Returns either the collection if available service addresses or error if something goes wrong.
	ResolveService(serviceName string) ([]string, error)
}
