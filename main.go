package main

import (
	"fmt"

	"github.com/agusheryanto182/go-online-store-mvp/config"
	rUser "github.com/agusheryanto182/go-online-store-mvp/domain/user/repository"
	sUser "github.com/agusheryanto182/go-online-store-mvp/domain/user/service"

	// hUser "github.com/agusheryanto182/go-online-store-mvp/domain/user/handler"

	hAuth "github.com/agusheryanto182/go-online-store-mvp/domain/auth/handler"
	rAuth "github.com/agusheryanto182/go-online-store-mvp/domain/auth/repository"
	sAuth "github.com/agusheryanto182/go-online-store-mvp/domain/auth/service"

	hProduct "github.com/agusheryanto182/go-online-store-mvp/domain/product/handler"
	rProduct "github.com/agusheryanto182/go-online-store-mvp/domain/product/repository"
	sProduct "github.com/agusheryanto182/go-online-store-mvp/domain/product/service"

	hCategory "github.com/agusheryanto182/go-online-store-mvp/domain/category/handler"
	rCategory "github.com/agusheryanto182/go-online-store-mvp/domain/category/repository"
	sCategory "github.com/agusheryanto182/go-online-store-mvp/domain/category/service"

	"github.com/agusheryanto182/go-online-store-mvp/helper/database"
	"github.com/agusheryanto182/go-online-store-mvp/helper/hashing"
	Njwt "github.com/agusheryanto182/go-online-store-mvp/helper/jwt"

	"github.com/agusheryanto182/go-online-store-mvp/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "Online Store MVP",
	})

	bootConfig := config.BootConfig()

	DB := database.InitialDB(*bootConfig)
	database.TableMigration(DB)
	hash := hashing.NewHash()
	jwt := Njwt.NewJWT(bootConfig.Secret)

	userRepo := rUser.NewUserRepository(DB)
	userService := sUser.NewUserService(userRepo)
	// userHandler := hUser.NewUserHandler(userService)

	authRepo := rAuth.NewAuthRepository(DB)
	authService := sAuth.NewAuthService(authRepo, userService, jwt, hash)
	authHandler := hAuth.NewAuthHandler(authService)

	categoryRepo := rCategory.NewCategoryRepository(DB)
	categoryService := sCategory.NewCategoryService(categoryRepo)
	categoryHandler := hCategory.NewCategoryHandler(categoryService)

	productRepo := rProduct.NewProductRepository(DB)
	productService := sProduct.NewProductService(productRepo, categoryService)
	productHandler := hProduct.NewProductService(productService)

	routes.AuthRoute(app, authHandler)
	routes.CategoryRoute(app, categoryHandler, jwt, userService)
	routes.ProductRoute(app, productHandler, jwt, userService)

	addr := fmt.Sprintf(":%d", bootConfig.AppPort)
	if err := app.Listen(addr).Error(); err != addr {
		panic("Appilaction failed to start")
	}
}
