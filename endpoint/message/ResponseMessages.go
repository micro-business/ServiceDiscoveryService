// Package message Defines all reply messages used in service discovery service
package message

// ResolveServiceResponse defines the message that contains the list of resolved services available in the network
type ResolveServiceResponse struct {
	ServiceAddresses []string `json:ServiceAddresses`
	Error            string   `json:"error,omitempty"`
}
