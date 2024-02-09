package auth

import (
	"github.com/agusheryanto182/go-online-store-mvp/domain/auth/dto"
	"github.com/agusheryanto182/go-online-store-mvp/entities"
	"github.com/gofiber/fiber/v2"
)

type AuthRepositoryInterface interface {
	InsertUser(newUser *entities.User) (*entities.User, error)
	FindUserByEmailOrUsername(identifier string) (*entities.User, error)
}

type AuthServiceInterface interface {
	Register(dataRequest *dto.RegisterRequest) (*entities.User, error)
	Login(dataRequest *dto.LoginRequest) (*entities.User, string, error)
}

type AuthHandlerInterface interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}
