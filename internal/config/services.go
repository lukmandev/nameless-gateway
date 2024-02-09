package config

import (
	"fmt"
	"os"
)

const (
	authServiceHostEnvName = "AUTH_SERVICE_HOST"
)

type ExternalServicesConfig interface {
	AuthServiceHost() string
}

type externalServiceConfig struct {
	authServiceHost string
}

func NewExternalServicesConfig() (ExternalServicesConfig, error) {
	authServiceHost := os.Getenv(authServiceHostEnvName)
	if len(authServiceHost) == 0 {
		return nil, fmt.Errorf("%s not found", authServiceHostEnvName)
	}

	return &externalServiceConfig{
		authServiceHost: authServiceHost,
	}, nil
}

func (cfg *externalServiceConfig) AuthServiceHost() string {
	return cfg.authServiceHost
}
