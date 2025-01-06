// By Saif Hamdan, Team Lead
// Date: 2025/1/6
package main

import (
	"log"

	"github.com/saifhamdan/go-apigateway-blueprint/config"
	"github.com/saifhamdan/go-apigateway-blueprint/pkg/db"
	"github.com/saifhamdan/go-apigateway-blueprint/pkg/http"
	"github.com/saifhamdan/go-apigateway-blueprint/pkg/logger"
	"github.com/saifhamdan/go-apigateway-blueprint/pkg/nats"

	"github.com/saifhamdan/go-apigateway-blueprint/internal/gateway"
)

const (
	ENV_PATH = "."
)

func main() {
	// Load configuration
	cfg, err := config.NewConfig(ENV_PATH)
	if err != nil {
		log.Fatalf("failed to load config from \"%s\"", ENV_PATH)
	}

	// Initialize logger
	logger := logger.NewLogger(cfg)

	// Flush logger before exit
	defer logger.Sync()

	// Initialize DB
	db := db.NewDB(cfg, logger)

	// Build postgres driver
	err = db.BuildPostgres()
	if err != nil {
		logger.Fatalf("failed to build db postgres driver")
	}
	defer db.Postgres.Close()

	// Build redis driver
	err = db.BuildRedis()
	if err != nil {
		logger.Fatalf("failed to build db redis driver")
	}
	defer db.Redis.Close()

	// Migrate Models related to postgres DB
	err = db.Postgres.MigrateDB()
	if err != nil {
		logger.Fatalf("failed to mirgare postgres db models")
	}

	// Initialize NATS
	nats, err := nats.NewNATS(cfg, logger)
	if err != nil {
		logger.Fatalf("failed to initialize NATS")
	}
	defer nats.Drain()
	defer nats.Close()

	// Initialize Fiber App
	app := http.NewApp(logger)

	// Start Gateway Server
	gateway := gateway.NewGateway(cfg, logger, db, nats, app)

	// Create done channel
	done := make(chan struct{})

	go gateway.Start(done)

	// Wait for the graceful shutdown to complete
	<-done

	logger.Info("Graceful shutdown complete.")
}
