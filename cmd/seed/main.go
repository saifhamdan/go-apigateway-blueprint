// By Saif Hamdan, Team Lead
// Date: 2025/1/6
package main

import (
	"log"

	"github.com/saifhamdan/go-apigateway-blueprint/config"
	"github.com/saifhamdan/go-apigateway-blueprint/data"
	"github.com/saifhamdan/go-apigateway-blueprint/pkg/db"
	"github.com/saifhamdan/go-apigateway-blueprint/pkg/logger"
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

	logger.Info("seeding data...")

	// Seed data
	data.Seed(logger, db)

	logger.Info("data seeded successfully")
}
