package routes

import (
	"github.com/agusheryanto182/go-online-store-mvp/domain/auth"
	"github.com/gofiber/fiber/v2"
)

func AuthRoute(app *fiber.App, handler auth.AuthHandlerInterface) {
	authGroup := app.Group("api/auth")
	authGroup.Post("/register", handler.Register)
	authGroup.Post("/login", handler.Login)
}
