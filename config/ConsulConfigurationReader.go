package config

import "github.com/microbusinesses/Micro-Businesses-Core/common/config"

// ConsulConfigurationReader implements ConfigurationReader using Consul to provide access to all configurations parameters required by the service.
type ConsulConfigurationReader struct {
	ConsulAddress              string
	ConsulScheme               string
	ListeningPortToOverride    int
	OverrideHostnameToOverride string
}

const serviceListeningPortKey = "services/service-discovery-service/endpoint/listening-port"
const overrideHostnameKey = "services/service-discovery-service/endpoint/override-hostname"

// GetListeningPort returns the port the application should start listening on.
// Returns either the listening port or error if something goes wrong.
func (consul ConsulConfigurationReader) GetListeningPort() (int, error) {
	if consul.ListeningPortToOverride != 0 {
		return consul.ListeningPortToOverride, nil

	}
	consulHelper := config.ConsulHelper{ConsulAddress: consul.ConsulAddress, ConsulScheme: consul.ConsulScheme}

	return consulHelper.GetInt(serviceListeningPortKey)
}

// GetOverrideHostname returns the overridden hostname to be used by the service discovery. If provided, the IP addresses or hostname
// provided by the service must be overridden by this parameter.
// This parameter is meant to be used in development environment.
func (consul ConsulConfigurationReader) GetOverrideHostname() (string, error) {
	if len(consul.OverrideHostnameToOverride) != 0 {
		return consul.OverrideHostnameToOverride, nil

	}

	consulHelper := config.ConsulHelper{ConsulAddress: consul.ConsulAddress, ConsulScheme: consul.ConsulScheme}

	return consulHelper.GetString(overrideHostnameKey)
}
