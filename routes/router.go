package routes

import (
	"github.com/agusheryanto182/go-online-store-mvp/domain/auth"
	"github.com/agusheryanto182/go-online-store-mvp/domain/category"
	"github.com/agusheryanto182/go-online-store-mvp/domain/product"
	"github.com/agusheryanto182/go-online-store-mvp/domain/user"
	"github.com/agusheryanto182/go-online-store-mvp/helper/jwt"
	"github.com/agusheryanto182/go-online-store-mvp/middleware"
	"github.com/gofiber/fiber/v2"
)

func AuthRoute(app *fiber.App, handler auth.AuthHandlerInterface) {
	authGroup := app.Group("api/auth")
	authGroup.Post("/register", handler.Register)
	authGroup.Post("/login", handler.Login)
}

func ProductRoute(app *fiber.App, handler product.ProductHandlerInterface, jwtService jwt.IJwt, userService user.UserServiceInterface) {
	productGroup := app.Group("api/product")
	productGroup.Post("/", middleware.Protected(jwtService, userService), handler.CreateProduct)
	productGroup.Put("/:id", middleware.Protected(jwtService, userService), handler.UpdateProduct)
	productGroup.Delete("/:id", middleware.Protected(jwtService, userService), handler.DeleteProduct)
	productGroup.Get("", handler.GetAllProducts)
	productGroup.Get("/:id", handler.GetProductByCategoryID)
}

func CategoryRoute(app *fiber.App, handler category.CategoryHandlerInterface, jwtService jwt.IJwt, userService user.UserServiceInterface) {
	categoryGroup := app.Group("api/category")
	categoryGroup.Post("/", middleware.Protected(jwtService, userService), handler.CreateCategory)
	categoryGroup.Patch("/:id", middleware.Protected(jwtService, userService), handler.UpdateCategory)
	categoryGroup.Delete("/:id", middleware.Protected(jwtService, userService), handler.DeleteCategory)
	categoryGroup.Get("", handler.GetAllCategory)
}
