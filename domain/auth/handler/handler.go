package handler

import (
	"github.com/agusheryanto182/go-online-store-mvp/domain/auth"
	"github.com/agusheryanto182/go-online-store-mvp/domain/auth/dto"
	"github.com/agusheryanto182/go-online-store-mvp/helper/response"
	"github.com/agusheryanto182/go-online-store-mvp/helper/validator"
	"github.com/gofiber/fiber/v2"
)

type AuthHandlerImpl struct {
	authService auth.AuthServiceInterface
}

func NewAuthHandler(authService auth.AuthServiceInterface) auth.AuthHandlerInterface {
	return &AuthHandlerImpl{authService: authService}
}

func (h *AuthHandlerImpl) Register(c *fiber.Ctx) error {
	var input dto.RegisterRequest
	if err := c.BodyParser(&input); err != nil {
		return response.SendStatusBadRequest(c, "Invalid input:"+err.Error())
	}

	if err := validator.ValidateStruct(input); err != nil {
		return response.SendStatusBadRequest(c, "error validating input"+err.Error())
	}

	newUser, err := h.authService.Register(&input)
	if err != nil {
		return response.SendStatusBadRequest(c, err.Error())
	}

	return response.SendStatusCreatedWithDataResponse(c, "register success", newUser)
}

func (h *AuthHandlerImpl) Login(c *fiber.Ctx) error {
	var input dto.LoginRequest

	if err := c.BodyParser(&input); err != nil {
		return response.SendStatusBadRequest(c, "Invalid input:"+err.Error())
	}

	if err := validator.ValidateStruct(input); err != nil {
		return response.SendStatusBadRequest(c, "error validating input"+err.Error())
	}

	userLogin, accessToken, err := h.authService.Login(&input)
	if err != nil {
		if err.Error() == "user is not found" {
			return response.SendStatusNotFound(c, "user is not found")
		}
		return response.SendStatusUnauthorized(c, "incorrect password")
	}
	return response.SendStatusOkWithDataResponse(c, "login success", dto.LoginResponse(userLogin, accessToken))
}
