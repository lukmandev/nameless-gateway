package main

import (
	"context"
	"log"

	"github.com/lukmandev/nameless/gateway/internal/app"
)

func main() {
	ctx := context.Background()
	_app, err := app.NewApp(ctx)

	if err != nil {
		log.Fatalf("Failed to create app: %s", err.Error())
	}

	err = _app.Run()
	if err != nil {
		log.Fatalf("Failed to start app: %s", err.Error())
	}
}
