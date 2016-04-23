package config

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/hashicorp/consul/api"
)

type ConsulConfigurationReader struct {
	ConsulAddress           string
	ConsulScheme            string
	ListeningPortToOverride int
}

const serviceListeningPortKey = "services/service-discovery/endpoint/listening-port"

func (consul ConsulConfigurationReader) GetListeningPort() (int, error) {
	if consul.ListeningPortToOverride != 0 {
		return consul.ListeningPortToOverride, nil

	} else {
		return getInt(consul, serviceListeningPortKey)
	}
}

func getKV(consul ConsulConfigurationReader) (*api.KV, error) {
	config := api.DefaultConfig()

	if len(consul.ConsulAddress) != 0 && len(consul.ConsulScheme) != 0 {
		config.Address = consul.ConsulAddress
		config.Scheme = consul.ConsulScheme
	}

	if client, err := api.NewClient(config); err != nil {
		return nil, err
	} else {
		return client.KV(), nil
	}
}

func getKeyPair(consul ConsulConfigurationReader, configKeyPath string) (*api.KVPair, error) {
	kv, err := getKV(consul)

	if err != nil {
		return nil, err
	}

	if keyPair, _, err := kv.Get(configKeyPath, nil); err != nil {
		return nil, err
	} else {
		return keyPair, nil
	}
}

func getInt(consul ConsulConfigurationReader, keyPath string) (int, error) {
	keyPair, err := getKeyPair(consul, keyPath)

	if err != nil {
		return 0, err
	}

	if keyPair == nil {
		return 0, errors.New(fmt.Sprintf("Consul key %s does not exist.", keyPath))

	}

	valueInString := string(keyPair.Value)

	if len(valueInString) == 0 {
		return 0, errors.New(fmt.Sprintf("Consul key %s is empty.", keyPath))

	}

	if value, err := strconv.Atoi(valueInString); err != nil {
		return 0, err
	} else {
		if value == 0 {
			return 0, errors.New(fmt.Sprintf("Consul key %s is zero.", keyPath))
		}

		return value, nil
	}
}
