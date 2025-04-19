package application

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shannon3335/story-server/internal/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	e      *echo.Echo
	server string
	DB     *gorm.DB
}

func New(server string) *App {
	ec := echo.New()
	db, err := connectDatabase()
	if err != nil {
		log.Fatal("Unable to connect to Database", err)
	}
	return &App{
		e:      ec,
		server: server,
		DB:     db,
	}
}

func connectDatabase() (*gorm.DB, error) {
	dsn := "host=localhost user=user password=pass dbname=story-server port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database %w", err)
	}

	// run migrations if the database connection isn't working
	if err := db.AutoMigrate(&types.StartStoryPrompt{}, &types.User{}); err != nil {
		return nil, fmt.Errorf("Failed to migrate database: %w", err)
	}
	return db, nil
}

func (app *App) Start(ctx context.Context) {
	fmt.Println("Starting server...")
	// Start the server in a goroutine
	go func() {
		if err := app.e.Start(":" + app.server); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Shutting down the server due to error: %v", err)
		}
	}()

	// Wait for termination signal
	<-ctx.Done()
	fmt.Println("Shutdown signal received, shutting down server...")

	// Gracefully shutdown the server
	if err := app.e.Shutdown(context.Background()); err != nil {
		log.Fatalf("Server shutdown failed:%+v", err)
	}
}
