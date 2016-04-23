package main

import (
	"flag"
	"os"
	"strconv"

	"github.com/microbusinesses/Micro-Businesses-Core/common/diagnostics"
	businessService "github.com/microbusinesses/ServiceDiscovery/business/service"
	"github.com/microbusinesses/ServiceDiscovery/config"
	"github.com/microbusinesses/ServiceDiscovery/endpoint"
)

var consulAddress string
var consulScheme string
var listeningPort int

func main() {
	flag.StringVar(&consulAddress, "consul-address", "", "The consul address in form of host:port. The default value is empty string.")
	flag.StringVar(&consulScheme, "consul-scheme", "", "The consul scheme. The default value is empty string.")
	flag.IntVar(&listeningPort, "listening-port", 0, "The port the application is serving HTTP request on. The default is zero.")
	flag.Parse()

	consulConfigurationReader := config.ConsulConfigurationReader{ConsulAddress: consulAddress, ConsulScheme: consulScheme}

	setConsulConfigurationValuesRequireToBeOverriden(&consulConfigurationReader)

	endpoint := endpoint.Endpoint{ConfigurationReader: consulConfigurationReader}

	serviceDiscoveryService := businessService.ConsulServiceDiscoveryService{}

	endpoint.ServiceDiscoveryService = serviceDiscoveryService

	endpoint.StartServer()
}

func setConsulConfigurationValuesRequireToBeOverriden(consulConfigurationReader *config.ConsulConfigurationReader) {
	diagnostics.IsNotNil(consulConfigurationReader, "consulConfigurationReader", "consulConfigurationReader is nil.")

	if listeningPort != 0 {
		consulConfigurationReader.ListeningPortToOverride = listeningPort
	} else if port, err := strconv.Atoi(os.Getenv("PORT")); err == nil && port != 0 {
		consulConfigurationReader.ListeningPortToOverride = port
	}
}
