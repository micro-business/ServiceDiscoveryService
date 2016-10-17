package endpoint

import (
	"errors"
	"strings"

	"github.com/go-kit/kit/endpoint"
	"github.com/graphql-go/graphql"
	"github.com/microbusinesses/ServiceDiscoveryService/business/contract"
	"golang.org/x/net/context"
)

const (
	address = "Address"
	port    = "Port"
)

type discoveredServiceInfo struct {
	Address string `json:"Address"`
	Port    int    `json:"Port"`
}

var discoveredServiceInfoType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Address",
		Fields: graphql.Fields{
			address: &graphql.Field{Type: graphql.String},
			port:    &graphql.Field{Type: graphql.String},
		},
	},
)

var rootQueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"servicesInfo": &graphql.Field{
				Type:        graphql.NewList(discoveredServiceInfoType),
				Description: "Returns the list of discovered services in the network",
				Args: graphql.FieldConfigArgument{
					"serviceName": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(resolveParams graphql.ResolveParams) (interface{}, error) {
					executionContext := resolveParams.Context.Value("ExecutionContext").(executionContext)
					serviceName, _ := resolveParams.Args["serviceName"].(string)

					var servicesInfo []contract.DiscoveredServiceInfo
					var err error

					if servicesInfo, err = executionContext.serviceDiscoveryService.ResolveService(
						serviceName); err != nil {
						return nil, err
					}

					result := make([]discoveredServiceInfo, 0, len(servicesInfo))

					for _, serviceInfo := range servicesInfo {
						result = append(result, discoveredServiceInfo{serviceInfo.Address, serviceInfo.Port})
					}

					return result, nil

				},
			},
		},
	},
)

var serviceDiscoveryServiceSchema, _ = graphql.NewSchema(graphql.SchemaConfig{Query: rootQueryType})

type executionContext struct {
	serviceDiscoveryService contract.ServiceDiscoveryService
}

func createAPIEndpoint(serviceDiscoveryService contract.ServiceDiscoveryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		result := executeQuery(request.(string), serviceDiscoveryService)

		if result.HasErrors() {
			errorMessages := []string{}

			for _, err := range result.Errors {
				errorMessages = append(errorMessages, err.Error())
			}

			return nil, errors.New(strings.Join(errorMessages, "\n"))
		}

		return result, nil
	}
}

func executeQuery(query string, serviceDiscoveryService contract.ServiceDiscoveryService) *graphql.Result {
	return graphql.Do(
		graphql.Params{
			Schema:        serviceDiscoveryServiceSchema,
			RequestString: query,
			Context:       context.WithValue(context.Background(), "ExecutionContext", executionContext{serviceDiscoveryService}),
		})
}
