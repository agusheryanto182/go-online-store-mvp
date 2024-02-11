package handler

import (
	"strconv"

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
		return response.SendStatusInternalServerError(c, "failed add the product : "+err.Error())
	}
	return response.SendStatusCreatedWithDataResponse(c, "success add the product", result)
}

func (h *CartHandlerImpl) GetCart(c *fiber.Ctx) error {
	currentUser, _ := c.Locals("CurrentUser").(*entities.User)
	if currentUser.Role != "user" {
		return response.SendStatusForbidden(c, "Access denied: you are admin, not user")
	}

	cartItemSummary, err := h.cartService.GetCart(currentUser.ID)
	if err != nil {
		return response.SendStatusOkResponse(c, "failed to get cart: "+err.Error())
	}

	return response.SendStatusOkWithDataResponse(c, "success to get cart", dto.FormatCart(cartItemSummary))
}

func (h *CartHandlerImpl) DeleteCartItem(c *fiber.Ctx) error {
	currentUser, _ := c.Locals("CurrentUser").(*entities.User)
	if currentUser.Role != "user" {
		return response.SendStatusForbidden(c, "Access denied: you are admin, not user")
	}
	id, err := strconv.Atoi(c.Params("product_id"))
	if err != nil {
		return response.SendStatusBadRequest(c, "invalid id")
	}
	err = h.cartService.RemoveProductFromCart(currentUser.ID, id)
	if err != nil {
		return response.SendStatusInternalServerError(c, "failed to delete product from cart : "+err.Error())
	}

	return response.SendStatusOkResponse(c, "success delete product from cart")

}
