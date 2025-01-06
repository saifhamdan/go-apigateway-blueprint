// By Saif Hamdan, Team Lead
// Date: 2025/1/6
package v1

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func (g *GatewayV1) RegisterGatewayV1Routes(root fiber.Router) {
	api := root.Group("/api/v1")
	userRoutes := api.Group("/users")

	// *************************  Middlewares *************************
	// your middlewares go here ex:
	// api.Use(middleware.AuthMiddleware)

	// ************************* Swagger Routes *************************
	root.Get("/swagger/v1/*", swagger.HandlerDefault)

	// ************************* User Routes *************************
	userRoutes.Get("/", g.GetUsers)
	userRoutes.Get("/:id", g.GetUser)
	userRoutes.Post("/", g.CreateUser)
	userRoutes.Patch("/:id", g.UpdateUser)
	userRoutes.Delete("/:id", g.DeleteUser)

	// ************************* Web Static Routes *************************
	// get html, css, js and images etc.....
	root.Static("/web", "public/web")
	// in case the URL was meant for the web application
	root.Get("/web/*", func(c *fiber.Ctx) error {
		return c.SendFile("public/web/index.html")
	})

	root.All("*", func(c *fiber.Ctx) error {
		return g.App.HttpResponseNotFound(c, fmt.Errorf("route not found 404"))
	})
}
