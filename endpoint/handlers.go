package endpoint

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/microbusinesses/Micro-Businesses-Core/common/diagnostics"
	businessContract "github.com/microbusinesses/ServiceDiscoveryService/business/contract"
	"github.com/microbusinesses/ServiceDiscoveryService/config"
	"github.com/microbusinesses/ServiceDiscoveryService/endpoint/transport"
	"golang.org/x/net/context"
)

type Endpoint struct {
	ConfigurationReader     config.ConfigurationReader
	ServiceDiscoveryService businessContract.ServiceDiscoveryService
}

func (endpoint Endpoint) StartServer() {
	diagnostics.IsNotNil(endpoint.ServiceDiscoveryService, "endpoint.ServiceDiscoveryService", "ServiceDiscoveryService must be provided.")
	diagnostics.IsNotNil(endpoint.ConfigurationReader, "endpoint.ConfigurationReader", "ConfigurationReader must be provided.")

	ctx := context.Background()

	if handlers, err := getHandlers(endpoint, ctx); err != nil {
		log.Fatal(err.Error())
	} else {
		http.HandleFunc("/CheckHealth", checkHealthHandleFunc)

		for pattern, handler := range handlers {
			http.Handle(pattern, handler)
		}

		if listeningPort, err := endpoint.ConfigurationReader.GetListeningPort(); err != nil {
			log.Fatal(err.Error())
		} else {
			log.Fatal(http.ListenAndServe(":"+strconv.Itoa(listeningPort), nil))
		}
	}
}

func getHandlers(endpoint Endpoint, ctx context.Context) (map[string]http.Handler, error) {
	handlers := make(map[string]http.Handler)

	if handler, err := createResolveServiceHandler(endpoint, ctx); err != nil {
		return map[string]http.Handler{}, err
	} else {
		handlers["/ResolveService"] = handler
	}

	return handlers, nil
}

func checkHealthHandleFunc(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Alive")
}

func createResolveServiceHandler(endpoint Endpoint, ctx context.Context) (http.Handler, error) {
	return httptransport.NewServer(
		ctx,
		createResolveServiceEndpoint(endpoint.ServiceDiscoveryService),
		transport.DecodeResolveServiceRequest,
		transport.EncodeResolveServiceResponse), nil
}
