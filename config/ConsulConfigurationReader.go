package config

import "github.com/microbusinesses/Micro-Businesses-Core/common/config"

type ConsulConfigurationReader struct {
	ConsulAddress           string
	ConsulScheme            string
	ListeningPortToOverride int
}

const serviceListeningPortKey = "services/service-discovery-service/endpoint/listening-port"

func (consul ConsulConfigurationReader) GetListeningPort() (int, error) {
	if consul.ListeningPortToOverride != 0 {
		return consul.ListeningPortToOverride, nil

	} else {
		consulHelper := config.ConsulHelper{ConsulAddress: consul.ConsulAddress, ConsulScheme: consul.ConsulScheme}

		return consulHelper.GetInt(serviceListeningPortKey)
	}
}
