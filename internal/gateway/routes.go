// By Saif Hamdan, Team Lead
// Date: 2025/1/6
package gateway

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"

	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"

	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recoverer "github.com/gofiber/fiber/v2/middleware/recover"
)

func (g *Gateway) RegisterGatewayRoutes() {
	root := g.App.Group("/")
	systemRoutes := root.Group("/system")

	// ************************* Global middlewares *************************
	// Import the middleware package that is part of the Fiber web framework
	root.Use(recoverer.New())

	// Import the middleware package that is part of the Fiber web framework
	root.Use(logger.New())

	// Limiter middleware for Fiber that is used to limit repeat requests to public APIs
	// and/or endpoints such as password reset. It is also useful for
	// API clients, web crawling, or other tasks that need to be throttled.
	root.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        100,
		Expiration: 30 * time.Second,
		LimitReached: func(c *fiber.Ctx) error {
			return g.App.HttpResponseTooManyRequests(c)
		},
	}))

	// Initialize Helmet middleware helps secure your apps by setting various HTTP headers.
	root.Use(helmet.New())
	// root.Use(helmet.New(helmet.Config{
	// 	XSSProtection:             "1; mode=block",
	// 	ContentTypeNosniff:        "nosniff",
	// 	XFrameOptions:             "DENY",
	// 	HSTSMaxAge:                31536000, // 1 year in seconds
	// 	HSTSExcludeSubdomains:     true,
	// 	ContentSecurityPolicy:     "default-src 'self'; script-src 'self'; object-src 'none';",
	// 	CSPReportOnly:             false,
	// 	HSTSPreloadEnabled:        true,
	// 	ReferrerPolicy:            "no-referrer",
	// 	PermissionPolicy:          "",
	// 	CrossOriginEmbedderPolicy: "require-corp",
	// 	CrossOriginOpenerPolicy:   "same-origin",
	// 	CrossOriginResourcePolicy: "same-origin",
	// 	OriginAgentCluster:        "?1",
	// 	XDNSPrefetchControl:       "off",
	// 	XDownloadOptions:          "noopen",
	// 	XPermittedCrossDomain:     "none",
	// }))

	// Encrypt Cookie is a middleware for Fiber that secures your cookie values through encryption.
	root.Use(encryptcookie.New(encryptcookie.Config{
		Key: g.Cfg.HttpCookieSecret,
	}))

	// Idempotency middleware for Fiber allows for fault-tolerant APIs
	// where duplicate requests — for example due to networking issues on the client-side
	// do not erroneously cause the same action performed multiple times on the server-side.
	// the client should set in the request header a unique X-Idempotency-Key in each request
	// in order to work
	root.Use(idempotency.New())

	// Compression middleware for Fiber that will compress the response using
	// gzip, deflate, brotli, and zstd compression depending on the Accept-Encoding header.
	root.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))

	// CORS (Cross-Origin Resource Sharing) is a middleware for Fiber that allows servers to specify who can access its resources and how
	root.Use(cors.New(cors.Config{
		AllowOrigins:  "http://localhost:5000, http://localhost:3000",
		AllowHeaders:  "Origin, Content-Type, Accept, Authorization",
		AllowMethods:  "GET, POST, PUT, DELETE",
		ExposeHeaders: "Content-Length",
		MaxAge:        3600,
	}))

	// ETag middleware for Fiber that lets caches be more efficient and save bandwidth
	// as a web server does not need to resend a full response if the content has not changed.
	root.Use(etag.New())

	// Initialize favicon to ignore favicon requests
	// and serve them by loading them in memory
	// root.Use(favicon.New(favicon.Config{
	// 	File: "./favicon.ico",
	// 	URL:  "/favicon.ico",
	// }))

	// Initialize healthcheck
	systemRoutes.Get("/live", healthcheck.New())

	// Monitor middleware for Fiber that reports server metrics
	systemRoutes.Get("/metrics", monitor.New())

	// Register v1 routes
	g.V1.RegisterGatewayV1Routes(root)
}
