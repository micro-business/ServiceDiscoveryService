package config

// ConfigurationReader defines the interface that provides access to all configurations parameters required by the service.
type ConfigurationReader interface {
	// GetListeningPort returns the port the application should start listening on.
	// Returns either the listening port or error if something goes wrong.
	GetListeningPort() (int, error)

	// GetOverrideHostname returns the overridden hostname to be used by the service discovery. If provided, the IP addresses or hostname
	// provided by the service must be overridden by this parameter.
	// This parameter is meant to be used in development environment.
	GetOverrideHostname() (string, error)
}
