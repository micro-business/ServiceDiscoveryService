package config

type ConfigurationReader interface {
	GetListeningPort() (int, error)
}
