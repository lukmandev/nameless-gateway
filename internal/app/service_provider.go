package app

import (
	"fmt"
	"log"

	"github.com/lukmandev/nameless-auth/pkg/auth_v1"
	"github.com/lukmandev/nameless-auth/pkg/user_v1"
	"github.com/lukmandev/nameless/gateway/internal/client"
	"github.com/lukmandev/nameless/gateway/internal/config"
	"github.com/lukmandev/nameless/gateway/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type serviceProvider struct {
	graphQLConfig config.GraphQLConfig

	serviceClients *client.ServiceClients

	authService service.AuthService
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

func (s *serviceProvider) ServiceClients() *client.ServiceClients {
	if s.serviceClients == nil {
		conn, err := grpc.Dial(
			fmt.Sprintf(":%d", 50051),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)

		if err != nil {
			log.Fatalf("Failed to connect to auth service: %s", err.Error())
		}

		fmt.Println("Connected to auth service")

		authClient := auth_v1.NewAuthV1Client(conn)
		userClient := user_v1.NewUserV1Client(conn)

		s.serviceClients = &client.ServiceClients{
			AuthClient: authClient,
			UserClient: userClient,
		}
	}

	return s.serviceClients
}
