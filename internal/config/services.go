package config

import (
	"fmt"
	"os"
)

const (
	authServiceHostEnvName  = "AUTH_SERVICE_HOST"
	movieServiceHostEnvName = "MOVIE_SERVICE_HOST"
)

type ExternalServicesConfig interface {
	AuthServiceHost() string
	MovieServiceHost() string
}

type externalServiceConfig struct {
	authServiceHost  string
	movieServiceHost string
}

func NewExternalServicesConfig() (ExternalServicesConfig, error) {
	authServiceHost := os.Getenv(authServiceHostEnvName)
	if len(authServiceHost) == 0 {
		return nil, fmt.Errorf("%s not found", authServiceHostEnvName)
	}

	movieServiceHost := os.Getenv(movieServiceHostEnvName)
	if len(movieServiceHost) == 0 {
		return nil, fmt.Errorf("%s not found", movieServiceHostEnvName)
	}

	return &externalServiceConfig{
		authServiceHost:  authServiceHost,
		movieServiceHost: movieServiceHost,
	}, nil
}

func (cfg *externalServiceConfig) AuthServiceHost() string {
	return cfg.authServiceHost
}

func (cfg *externalServiceConfig) MovieServiceHost() string {
	return cfg.movieServiceHost
}
