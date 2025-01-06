// By Saif Hamdan, Team Lead
// Date: 2025/1/6
package gateway

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/saifhamdan/go-apigateway-blueprint/config"
	"github.com/saifhamdan/go-apigateway-blueprint/internal/gateway/middlewares"
	v1 "github.com/saifhamdan/go-apigateway-blueprint/internal/gateway/v1"
	"github.com/saifhamdan/go-apigateway-blueprint/pkg/db"
	"github.com/saifhamdan/go-apigateway-blueprint/pkg/http"
	"github.com/saifhamdan/go-apigateway-blueprint/pkg/logger"
	"github.com/saifhamdan/go-apigateway-blueprint/pkg/nats"
	"github.com/saifhamdan/go-apigateway-blueprint/pkg/validator"
)

type Gateway struct {
	Cfg         *config.Config
	Logger      *logger.Logger
	DB          *db.DB
	Nats        *nats.Nats
	App         *http.App
	Validate    *validator.Validate
	V1          *v1.GatewayV1
	MiddleWares *middlewares.Middleware
}

func NewGateway(cfg *config.Config, logger *logger.Logger, db *db.DB, nats *nats.Nats, app *http.App) *Gateway {
	// Initialize Validator
	validator := validator.New()

	// Initialize middlewares
	middlewares := middlewares.NewMiddleware(cfg, logger, db, nats, app)

	// Initialize Gateway V1
	v1 := v1.NewGatewayV1(cfg, logger, db, nats, app, validator, middlewares)

	return &Gateway{
		Cfg:         cfg,
		Logger:      logger,
		DB:          db,
		Nats:        nats,
		App:         app,
		Validate:    validator,
		V1:          v1,
		MiddleWares: middlewares,
	}
}

func (g *Gateway) Start(done chan struct{}) {
	url := fmt.Sprintf("%s:%s", g.Cfg.HttpHost, g.Cfg.HttpPort)

	// Register Routes
	g.RegisterGatewayRoutes()

	g.Logger.Infof("go-apigateway-blueprint server has started on: %s", url)
	go func() {
		if err := g.App.Listen(url); err != nil {
			g.Logger.Fatalf("failed to serve: %v", err)
		}
	}()

	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	g.Logger.Info("shutting down gracefully, press Ctrl+C again to force")

	g.Logger.Info("shutting down gateway server...")

	// Shutdown Fiber App Gateway gracefully
	g.App.Shutdown()

	g.Logger.Info("gateway server has gracefully stopped")

	// Notify the main goroutine that the shutdown is complete
	done <- struct{}{}
}
