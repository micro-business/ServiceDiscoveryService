package config

import "github.com/microbusinesses/Micro-Businesses-Core/common/config"

type ConsulConfigurationReader struct {
	ConsulAddress              string
	ConsulScheme               string
	ListeningPortToOverride    int
	OverrideHostnameToOverride string
}

const serviceListeningPortKey = "services/service-discovery-service/endpoint/listening-port"
const overrideHostnameKey = "services/service-discovery-service/endpoint/override-hostname"

func (consul ConsulConfigurationReader) GetListeningPort() (int, error) {
	if consul.ListeningPortToOverride != 0 {
		return consul.ListeningPortToOverride, nil

	} else {
		consulHelper := config.ConsulHelper{ConsulAddress: consul.ConsulAddress, ConsulScheme: consul.ConsulScheme}

		return consulHelper.GetInt(serviceListeningPortKey)
	}
}

func (consul ConsulConfigurationReader) GetOverrideHostname() (string, error) {
	if len(consul.OverrideHostnameToOverride) != 0 {
		return consul.OverrideHostnameToOverride, nil

	} else {
		consulHelper := config.ConsulHelper{ConsulAddress: consul.ConsulAddress, ConsulScheme: consul.ConsulScheme}

		return consulHelper.GetString(overrideHostnameKey)
	}
}
