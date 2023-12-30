package config

import (
	"errors"
	"net"
	"os"
)

const (
	graphQLHostEnvName       = "GRAPHQL_HOST"
	graphQLPortEnvName       = "GRAPHQL_PORT"
	graphQLPlaygroundEnvName = "GRAPHQL_PLAYGROUND"
)

type GraphQLConfig interface {
	Address() string
	Host() string
	Port() string
	PlaygroundEnabled() bool
}

type graphQLConfig struct {
	host              string
	port              string
	playgroundEnabled bool
}

func NewGraphQLConfig() (GraphQLConfig, error) {
	host := os.Getenv(graphQLHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("GRAPHQL host not found")
	}
	port := os.Getenv(graphQLPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("GRAPHQL port not found")
	}
	playgroundEnabled := os.Getenv(graphQLPlaygroundEnvName)
	if len(port) == 0 {
		return nil, errors.New("GRAPHQL port not found")
	}

	return &graphQLConfig{
		host:              host,
		port:              port,
		playgroundEnabled: playgroundEnabled == "true",
	}, nil
}

func (cfg *graphQLConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}

func (cfg *graphQLConfig) Host() string {
	return cfg.host
}

func (cfg *graphQLConfig) Port() string {
	return cfg.port
}

func (cfg *graphQLConfig) PlaygroundEnabled() bool {
	return cfg.playgroundEnabled
}
