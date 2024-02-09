package handler

import (
	"github.com/agusheryanto182/go-online-store-mvp/domain/user"
	"github.com/agusheryanto182/go-online-store-mvp/entities"
	"github.com/gofiber/fiber/v2"
)

type UserHandlerImpl struct {
	userService user.UserServiceInterface
}

func NewUserHandler(userService user.UserServiceInterface) user.UserHandlerInterface {
	return &UserHandlerImpl{userService: userService}
}

func (h UserHandlerImpl) GetCurrentUser(c *fiber.Ctx) {
	user, ok := c.Locals("CurrentUser").(*entities.User)
	if !ok || user == nil {
		return
	}
}
