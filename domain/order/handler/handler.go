package handler

import (
	"strconv"

	"github.com/agusheryanto182/go-online-store-mvp/domain/order"
	"github.com/agusheryanto182/go-online-store-mvp/domain/order/dto"
	"github.com/agusheryanto182/go-online-store-mvp/entities"
	"github.com/agusheryanto182/go-online-store-mvp/helper/response"
	"github.com/gofiber/fiber/v2"
)

type OrderHandlerImpl struct {
	orderService order.OrderServiceInterface
}

func NewOrderHandler(orderService order.OrderServiceInterface) order.OrderHandlerInterface {
	return &OrderHandlerImpl{orderService: orderService}
}

func (h *OrderHandlerImpl) CreateOrderFromProduct(c *fiber.Ctx) error {
	var req *dto.CreateOrderRequest

	currentUser := c.Locals("CurrentUser").(*entities.User)
	if currentUser.Role != "user" {
		return response.SendStatusForbidden(c, "Access denied: you are admin, not user")
	}

	if err := c.BodyParser(&req); err != nil {
		return response.SendStatusBadRequest(c, "invalid input : "+err.Error())
	}

	req.User = currentUser

	newOrder, err := h.orderService.CreateOrderFromProduct(currentUser.ID, req)
	if err != nil {
		return response.SendStatusBadRequest(c, "failed to create order : "+err.Error())
	}

	return response.SendStatusCreatedWithDataResponse(c, "success", dto.FormatOrderResponse(newOrder))
}

func (h *OrderHandlerImpl) CreateOrderFromCart(c *fiber.Ctx) error {

	currentUser := c.Locals("CurrentUser").(*entities.User)
	if currentUser.Role != "user" {
		return response.SendStatusForbidden(c, "Access denied: you are admin, not user")
	}

	id, err := strconv.Atoi(c.Params("cart_id"))
	if err != nil {
		return response.SendStatusBadRequest(c, "invalid id : "+err.Error())
	}

	req := &dto.CreateOrderRequestFromCart{
		CartID: id,
		User:   currentUser,
	}

	newOrder, err := h.orderService.CreateOrderFromCart(currentUser.ID, req)
	if err != nil {
		return response.SendStatusBadRequest(c, "failed to create order : "+err.Error())
	}

	return response.SendStatusCreatedWithDataResponse(c, "success", dto.FormatOrderResponse(newOrder))
}
