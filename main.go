package main

import (
	"flag"
	"os"
	"strconv"

	"github.com/micro-business/Micro-Business-Core/common/diagnostics"
	"github.com/micro-business/ServiceDiscoveryService/business/service"
	"github.com/micro-business/ServiceDiscoveryService/config"
	"github.com/micro-business/ServiceDiscoveryService/endpoint"
)

var consulAddress string
var consulScheme string
var listeningPort int
var overrideHostname string

func main() {
	flag.StringVar(&consulAddress, "consul-address", "", "The consul address in form of host:port. The default value is empty string.")
	flag.StringVar(&consulScheme, "consul-scheme", "", "The consul scheme. The default value is empty string.")
	flag.IntVar(&listeningPort, "listening-port", 0, "The port the application is serving HTTP request on. The default is zero.")
	flag.StringVar(&overrideHostname, "override-hostname", "", "The override host name, if provided, all returned IP addresses will be replaced by this host name. The default value is empty string.")
	flag.Parse()

	consulConfigurationReader := config.ConsulConfigurationReader{ConsulAddress: consulAddress, ConsulScheme: consulScheme}

	setConsulConfigurationValuesRequireToBeOverriden(&consulConfigurationReader)

	endpoint := endpoint.Endpoint{ConfigurationReader: consulConfigurationReader}

	serviceDiscoveryService := service.ConsulServiceDiscoveryService{ConsulAddress: consulAddress, ConsulScheme: consulScheme, ConfigurationReader: consulConfigurationReader}

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
