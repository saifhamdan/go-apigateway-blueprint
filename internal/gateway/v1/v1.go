// By Saif Hamdan, Team Lead
// Date: 2025/1/6
package v1

import (
	"github.com/saifhamdan/go-apigateway-blueprint/config"
	"github.com/saifhamdan/go-apigateway-blueprint/internal/gateway/middlewares"
	_ "github.com/saifhamdan/go-apigateway-blueprint/models/v1"
	"github.com/saifhamdan/go-apigateway-blueprint/pkg/db"
	"github.com/saifhamdan/go-apigateway-blueprint/pkg/http"
	"github.com/saifhamdan/go-apigateway-blueprint/pkg/logger"
	"github.com/saifhamdan/go-apigateway-blueprint/pkg/nats"
	"github.com/saifhamdan/go-apigateway-blueprint/pkg/validator"
	_ "github.com/saifhamdan/go-apigateway-blueprint/swagger/v1"
)

// @title						Gateway API Blueprint
// @version					1.0
// @description				Gateway APIs Docs for API Blueprint made by Saif Hamdan
// @termsOfService				https://saifhamdan14@gmail.com
// @contact.name				API Support
// @contact.email				support@saifhamdan14gmail.com
// @license.url				https://saifhamdan14@gmail.com
// @license.name				Generic Portal
// @host						localhost:8888
// @BasePath					/swagger/v1
// @Schemes					http https
type GatewayV1 struct {
	Cfg         *config.Config
	Logger      *logger.Logger
	DB          *db.DB
	Nats        *nats.Nats
	App         *http.App
	Validate    *validator.Validate
	Middlewares *middlewares.Middleware
}

func NewGatewayV1(cfg *config.Config, logger *logger.Logger, db *db.DB, nats *nats.Nats, app *http.App, v *validator.Validate, m *middlewares.Middleware) *GatewayV1 {
	return &GatewayV1{
		Cfg:      cfg,
		Logger:   logger,
		DB:       db,
		Nats:     nats,
		App:      app,
		Validate: v,
	}
}
