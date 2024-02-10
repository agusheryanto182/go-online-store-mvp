package handler

import (
	"github.com/agusheryanto182/go-online-store-mvp/domain/cart"
	"github.com/agusheryanto182/go-online-store-mvp/domain/cart/dto"
	"github.com/agusheryanto182/go-online-store-mvp/entities"
	"github.com/agusheryanto182/go-online-store-mvp/helper/response"
	"github.com/agusheryanto182/go-online-store-mvp/helper/validator"
	"github.com/gofiber/fiber/v2"
)

type CartHandlerImpl struct {
	cartService cart.CartServiceInterface
}

func NewCartHandler(cartService cart.CartServiceInterface) cart.CartHandlerInterface {
	return &CartHandlerImpl{cartService: cartService}
}

func (h *CartHandlerImpl) AddCartItem(c *fiber.Ctx) error {
	currentUser, _ := c.Locals("CurrentUser").(*entities.User)
	if currentUser.Role != "user" {
		return response.SendStatusForbidden(c, "Access denied: you are admin, not user")
	}

	req := new(dto.AddCartItemsRequest)
	if err := c.BodyParser(req); err != nil {
		return response.SendStatusBadRequest(c, "invalid payload:"+err.Error())
	}
	if err := validator.ValidateStruct(req); err != nil {
		return response.SendStatusBadRequest(c, "error validating payload:"+err.Error())
	}

	result, err := h.cartService.AddCartItems(currentUser.ID, req)
	if err != nil {
		return response.SendStatusInternalServerError(c, "Gagal menambahkan produk ke keranjang: "+err.Error())
	}
	return response.SendStatusCreatedWithDataResponse(c, "Berhasil menambahkan produk ke keranjang", result)
}
