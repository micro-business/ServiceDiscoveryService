package config

type ConfigurationReader interface {
	GetListeningPort() (int, error)

	GetOverrideHostname() (string, error)
}
