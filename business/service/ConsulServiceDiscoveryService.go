package service

import (
	"strconv"

	"github.com/hashicorp/consul/api"
	"github.com/microbusinesses/Micro-Businesses-Core/common/diagnostics"
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
// Returns either the collection if available service addresses or error if something goes wrong.
func (consulServiceDiscoveryService ConsulServiceDiscoveryService) ResolveService(serviceName string) ([]string, error) {
	diagnostics.IsNotNil(consulServiceDiscoveryService.ConfigurationReader, "consulServiceDiscoveryService.ConfigurationReader", "ConfigurationReader must be provided.")
	diagnostics.IsNotNilOrEmptyOrWhitespace(serviceName, "serviceName", "serviceName must be provided. Cannot be nil, empty or contains whitespace character only.")

	var serviceAddresses []string

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

	serviceAddresses = make([]string, len(checks))
	overrideHostname, _ := consulServiceDiscoveryService.ConfigurationReader.GetOverrideHostname()

	for idx, check := range checks {
		if len(overrideHostname) == 0 {
			serviceAddresses[idx] = check.Service.Address + ":" + strconv.Itoa(check.Service.Port)
		} else {
			serviceAddresses[idx] = overrideHostname + ":" + strconv.Itoa(check.Service.Port)
		}
	}

	return serviceAddresses, nil
}
