package service

import (
	"errors"
	"time"

	"github.com/agusheryanto182/go-online-store-mvp/domain/order"
	"github.com/agusheryanto182/go-online-store-mvp/domain/order/dto"
	payment "github.com/agusheryanto182/go-online-store-mvp/domain/payment/service"
	"github.com/agusheryanto182/go-online-store-mvp/domain/product"
	"github.com/agusheryanto182/go-online-store-mvp/entities"
)

type OrderServiceImpl struct {
	orderRepository order.OrderRepositoryInterface
	productService  product.ProductServiceInterface
	paymentService  payment.Service
}

func NewOrderService(
	orderRepository order.OrderRepositoryInterface,
	productService product.ProductServiceInterface,
	paymentService payment.Service,
) order.OrderServiceInterface {
	return &OrderServiceImpl{
		orderRepository: orderRepository,
		productService:  productService,
		paymentService:  paymentService,
	}
}

func (s *OrderServiceImpl) CreateOrderFromCart(userID int, request *dto.CreateOrderRequest) (*entities.Order, error) {
	order := &entities.Order{}
	uniqueID := time.Now().Format("20060102150405")

	getProduct, _ := s.productService.GetProductByID(request.ProductID)

	order.IdOrder = uniqueID
	order.UserId = userID
	order.GrandTotalQuantity = request.Quantity
	order.OrderStatus = "pending"
	order.TotalAmountPaid = getProduct.Price * request.Quantity

	newOrder, err := s.orderRepository.CreateOrder(order)
	if err != nil {
		return nil, errors.New("failed to create order")
	}

	paymentOrder := entities.Transaction{
		ID:     newOrder.Id,
		Amount: newOrder.TotalAmountPaid,
	}

	paymentURL, err := s.paymentService.GetPaymentURL(paymentOrder, *request.User)
	if err != nil {
		return nil, errors.New("failed to create payment url")
	}

	newOrder.PaymentURL = paymentURL

	newResult, err := s.orderRepository.UpdateOrder(newOrder.Id, newOrder)
	if err != nil {
		return nil, errors.New("failed to update order")
	}

	return newResult, nil

}
