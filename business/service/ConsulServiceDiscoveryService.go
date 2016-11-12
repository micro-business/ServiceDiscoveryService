package service

import (
	"github.com/hashicorp/consul/api"
	"github.com/micro-business/Micro-Business-Core/common/diagnostics"
	"github.com/micro-business/ServiceDiscoveryService/business/contract"
	"github.com/micro-business/ServiceDiscoveryService/config"
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

	serviceEntries, _, err := client.Health().Service(serviceName, "", true, nil)

	if err != nil {
		return nil, err
	}

	result := make([]contract.DiscoveredServiceInfo, 0, len(serviceEntries))
	overrideHostname, _ := consulServiceDiscoveryService.ConfigurationReader.GetOverrideHostname()

	for _, serviceEntry := range serviceEntries {
		var serviceInfo contract.DiscoveredServiceInfo

		if len(overrideHostname) == 0 {
			serviceInfo = contract.DiscoveredServiceInfo{serviceEntry.Service.Address, serviceEntry.Service.Port}
		} else {
			serviceInfo = contract.DiscoveredServiceInfo{overrideHostname, serviceEntry.Service.Port}
		}

		// Ignoring duplicated result
		if !isServiceInfoAlreadyInList(result, serviceInfo) {
			result = append(result, serviceInfo)
		}
	}

	return result, nil
}

func isServiceInfoAlreadyInList(servicesInfo []contract.DiscoveredServiceInfo, serviceInfo contract.DiscoveredServiceInfo) bool {
	for _, info := range servicesInfo {
		if info == serviceInfo {
			return true
		}
	}

	return false
}
