package main

import (
	"github.com/a1exCross/auth/internal/app"

	"context"
	"log"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %v", err)
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %v", err)
	}
}
