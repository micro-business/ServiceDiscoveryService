// Defines all request message used in service discovery service
package message

// ResolveServiceRequest defines the message that is used to resolve available services in the network
type ResolveServiceRequest struct {
	ServiceName string `json:ServiceName`
}
