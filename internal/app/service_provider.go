package app

import (
	"log"

	"github.com/lukmandev/nameless/gateway/internal/config"
)

type serviceProvider struct {
	graphQLConfig config.GraphQLConfig
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GraphQLConfig() config.GraphQLConfig {
	if s.graphQLConfig == nil {
		graphQLConfig, err := config.NewGraphQLConfig()
		if err != nil {
			log.Fatalf("Failed to create a graphql config: %s", err.Error())
		}
		s.graphQLConfig = graphQLConfig
	}

	return s.graphQLConfig
}
