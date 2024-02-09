package app

import (
	"context"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/lukmandev/nameless/gateway/internal/api"
	"github.com/lukmandev/nameless/gateway/internal/config"
	"github.com/lukmandev/nameless/gateway/internal/middleware"
)

type App struct {
	serviceProvider *serviceProvider
	graphQLServer   http.Handler
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
	resolver := api.NewResolver(a.serviceProvider.AuthService(ctx), a.serviceProvider.UserService(ctx))

	srv := handler.NewDefaultServer(api.NewExecutableSchema(api.Config{Resolvers: &resolver}))

	a.serviceProvider.AuthService(ctx)

	router := chi.NewRouter()

	router.Use(middleware.AuthMiddleware())

	playgroundEnabled := a.serviceProvider.GraphQLConfig().PlaygroundEnabled()

	if playgroundEnabled {
		router.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	}

	router.Handle("/graphql", srv)

	a.graphQLServer = router

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

func (a *App) initServiceClients(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) runGraphQL() error {
	host := a.serviceProvider.GraphQLConfig().Host()
	port := a.serviceProvider.GraphQLConfig().Port()

	router := a.graphQLServer

	log.Printf("connect to http://%s:%s/ for GraphQL playground", host, port)
	err := http.ListenAndServe(":"+port, router)

	return err
}

func (a *App) Run() error {

	return a.runGraphQL()
}
