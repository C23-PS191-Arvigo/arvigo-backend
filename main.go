package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/yusufwib/arvigo-backend/pkg/database"
	"github.com/yusufwib/arvigo-backend/route"
)

const (
	version = "1.0.5" // release-arvigo-backend-1.0.5
	appName = "arvigo-backend"
)

func main() {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to the database and run migrations
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatal("Error closing database")
		}
		sqlDB.Close()
	}()

	// Create a new Echo instance
	e := echo.New()

	// Add middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Register routes
	route.RegisterAllRoutes(e)

	// Start the server in a separate goroutine
	go func() {
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}

		log.Printf("Server listening on port %s. \nThis %s service's is using version %s", port, appName, version)
		err = e.Start(":" + port)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Set a timeout for the shutdown process
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Gracefully shutdown the server
	err = e.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server gracefully stopped")
}
