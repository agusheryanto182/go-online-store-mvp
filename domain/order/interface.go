package order

import (
	"github.com/agusheryanto182/go-online-store-mvp/domain/order/dto"
	"github.com/agusheryanto182/go-online-store-mvp/entities"
	"github.com/gofiber/fiber/v2"
)

type OrderRepositoryInterface interface {
	CreateOrder(newOrder *entities.Order) (*entities.Order, error)
	UpdateOrder(ID int, updateOrder *entities.Order) (*entities.Order, error)
}

type OrderServiceInterface interface {
	CreateOrderFromProduct(userID int, request *dto.CreateOrderRequest) (*entities.Order, error)
	CreateOrderFromCart(userID int, request *dto.CreateOrderRequestFromCart) (*entities.Order, error)
}

type OrderHandlerInterface interface {
	CreateOrderFromProduct(c *fiber.Ctx) error
	CreateOrderFromCart(c *fiber.Ctx) error
}
