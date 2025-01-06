// By Saif Hamdan, Team Lead
// Date: 2025/1/6
package db

import (
	"github.com/saifhamdan/go-apigateway-blueprint/pkg/logger"

	"github.com/saifhamdan/go-apigateway-blueprint/config"
)

type DB struct {
	cfg      *config.Config
	logger   *logger.Logger
	Postgres *postgres
	Redis    *redis
}

func NewDB(cfg *config.Config, logger *logger.Logger) *DB {
	return &DB{
		cfg:    cfg,
		logger: logger,
	}
}

func (d *DB) BuildPostgres() error {
	postgres, err := d.newPostgres()
	if err != nil {
		return err
	}

	d.Postgres = postgres

	return nil
}

func (d *DB) BuildRedis() error {
	redis, err := d.newRedis()
	if err != nil {
		return err
	}

	d.Redis = redis

	return nil
}
