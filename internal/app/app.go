package app

import (
	"context"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/lukmandev/nameless/gateway/internal/api"
	"github.com/lukmandev/nameless/gateway/internal/config"
)

const defaultPort = "4000"

type App struct {
	serviceProvider *serviceProvider
	graphQLServer   *handler.Server
}

func NewApp(ctx context.Context) (*App, error) {
	_app := &App{}

	err := _app.initDeps(ctx)
	if err != nil {
		log.Fatalf("Failed to init dependencies: %s", err.Error())
	}

	return _app, nil
}

func (a *App) initConfig(ctx context.Context) error {
	config.Load(".env")

	return nil
}

func (a *App) initGraphQL(ctx context.Context) error {
	resolver := api.NewResolver()

	srv := handler.NewDefaultServer(api.NewExecutableSchema(api.Config{Resolvers: &resolver}))

	a.graphQLServer = srv

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGraphQL,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) runGraphQL() error {
	host := a.serviceProvider.GraphQLConfig().Host()
	port := a.serviceProvider.GraphQLConfig().Port()
	playgroundEnabled := a.serviceProvider.GraphQLConfig().PlaygroundEnabled()

	if playgroundEnabled {
		http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	}

	http.Handle("/graphql", a.graphQLServer)

	log.Printf("connect to http://%s:%s/ for GraphQL playground", host, port)
	err := http.ListenAndServe(":"+port, nil)

	return err
}

func (a *App) Run() error {

	return a.runGraphQL()
}
