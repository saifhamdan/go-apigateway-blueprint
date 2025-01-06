// By Saif Hamdan, Team Lead
// Date: 2025/1/6
package middlewares

import (
	"github.com/saifhamdan/go-apigateway-blueprint/config"
	"github.com/saifhamdan/go-apigateway-blueprint/pkg/db"
	"github.com/saifhamdan/go-apigateway-blueprint/pkg/http"
	"github.com/saifhamdan/go-apigateway-blueprint/pkg/logger"
	"github.com/saifhamdan/go-apigateway-blueprint/pkg/nats"
)

type Middleware struct {
	Cfg    *config.Config
	Logger *logger.Logger
	DB     *db.DB
	Nats   *nats.Nats
	App    *http.App
}

func NewMiddleware(cfg *config.Config, logger *logger.Logger, db *db.DB, nats *nats.Nats, app *http.App) *Middleware {

	m := &Middleware{
		Cfg:    cfg,
		App:    app,
		Logger: logger,
		DB:     db,
		Nats:   nats,
	}

	return m
}
