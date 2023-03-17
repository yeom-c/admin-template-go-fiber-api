package route

import (
	"github.com/yeom-c/admin-template-go-fiber-api/app"
	"github.com/yeom-c/admin-template-go-fiber-api/handler"
)

func SetRoutes() {
	h := handler.Handler()

	app.Server().Fiber.Get("/health", h.Health)
	app.Server().Fiber.Post("/sign-in", h.SignIn)
}
