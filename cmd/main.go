package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"github.com/shannon3335/story-server/application"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env files", err)
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		log.Fatal("SERVER_PORT is not set")
	}

	app := application.New(port)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	app.LoadRoutes()
	app.Start(ctx)
}
