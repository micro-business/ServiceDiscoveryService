package service

import (
	"github.com/hashicorp/consul/api"
	"github.com/microbusinesses/Micro-Businesses-Core/common/diagnostics"
	"github.com/microbusinesses/ServiceDiscoveryService/business/contract"
	"github.com/microbusinesses/ServiceDiscoveryService/config"
)

// ConsulServiceDiscoveryService uses Consul to resolve available  services in the network.
type ConsulServiceDiscoveryService struct {
	ConsulAddress       string
	ConsulScheme        string
	ConfigurationReader config.ConfigurationReader
}

// ResolveService resolves the provided service name by returning the list of services providing the same functionality.
// serviceName: Mandatory. The name of the service to resolve
// Returns either the collection of available service information or error if something goes wrong.
func (consulServiceDiscoveryService ConsulServiceDiscoveryService) ResolveService(serviceName string) ([]contract.DiscoveredServiceInfo, error) {
	diagnostics.IsNotNil(consulServiceDiscoveryService.ConfigurationReader, "consulServiceDiscoveryService.ConfigurationReader", "ConfigurationReader must be provided.")
	diagnostics.IsNotNilOrEmptyOrWhitespace(serviceName, "serviceName", "serviceName must be provided. Cannot be nil, empty or contains whitespace character only.")

	config := api.DefaultConfig()

	if len(consulServiceDiscoveryService.ConsulAddress) != 0 && len(consulServiceDiscoveryService.ConsulScheme) != 0 {
		config.Address = consulServiceDiscoveryService.ConsulAddress
		config.Scheme = consulServiceDiscoveryService.ConsulScheme
	}

	client, err := api.NewClient(config)

	if err != nil {
		return nil, err
	}

	checks, _, err := client.Health().Service(serviceName, "", true, nil)

	if err != nil {
		return nil, err
	}

	discoveredServicesInfo := make([]contract.DiscoveredServiceInfo, len(checks))
	overrideHostname, _ := consulServiceDiscoveryService.ConfigurationReader.GetOverrideHostname()

	for _, check := range checks {
		if len(overrideHostname) == 0 {
			discoveredServicesInfo = append(discoveredServicesInfo, contract.DiscoveredServiceInfo{check.Service.Address, check.Service.Port})
		} else {
			discoveredServicesInfo = append(discoveredServicesInfo, contract.DiscoveredServiceInfo{overrideHostname, check.Service.Port})
		}
	}

	return discoveredServicesInfo, nil
}
