// Package contract defines the service discovery contract.
package contract

// DiscoveredServiceInfo defines the information about the discovered service in the network.
type DiscoveredServiceInfo struct {
	// The hostname or IP address of the discovered service.
	Address string

	// The listening port of the discovered service.
	Port int
}

// ServiceDiscoveryService contract defines the service that can resolve addresses for the available services in the network.
type ServiceDiscoveryService interface {
	// ResolveService resolves the provided service name by returning the list of services providing the same functionality.
	// serviceName: Mandatory. The name of the service to resolve
	// Returns either the collection of available service information or error if something goes wrong.
	ResolveService(serviceName string) ([]DiscoveredServiceInfo, error)
}
