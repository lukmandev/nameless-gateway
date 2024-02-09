package app

import (
	"fmt"
	"log"

	authDesc "github.com/lukmandev/nameless-auth/pkg/auth_v1"
	userDesc "github.com/lukmandev/nameless-auth/pkg/user_v1"
	"github.com/lukmandev/nameless/gateway/internal/client"
	authClient "github.com/lukmandev/nameless/gateway/internal/client/auth"
	userClient "github.com/lukmandev/nameless/gateway/internal/client/user"
	"github.com/lukmandev/nameless/gateway/internal/config"
	"github.com/lukmandev/nameless/gateway/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type serviceProvider struct {
	graphQLConfig          config.GraphQLConfig
	externalServicesConfig config.ExternalServicesConfig

	serviceClients *client.ServiceClients

	authService service.AuthService
	userService service.UserService
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

func (s *serviceProvider) ExternalServicesConfig() config.ExternalServicesConfig {
	if s.externalServicesConfig == nil {
		conf, err := config.NewExternalServicesConfig()
		if err != nil {
			log.Fatalf("Failed to create a external services config: %s", err.Error())
		}
		s.externalServicesConfig = conf
	}

	return s.externalServicesConfig
}

func (s *serviceProvider) ServiceClients() *client.ServiceClients {
	if s.serviceClients == nil {
		conn, err := grpc.Dial(
			s.ExternalServicesConfig().AuthServiceHost(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)

		if err != nil {
			log.Fatalf("Failed to connect to auth service: %s", err.Error())
		}

		fmt.Println("Connected to auth service")

		authServiceClient := authDesc.NewAuthV1Client(conn)
		userServiceClient := userDesc.NewUserV1Client(conn)

		newAuthClient := authClient.New(authServiceClient)
		newUserClient := userClient.New(userServiceClient)

		s.serviceClients = client.New(newAuthClient, newUserClient)
	}

	return s.serviceClients
}
