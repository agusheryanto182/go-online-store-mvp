package dto

import "github.com/agusheryanto182/go-online-store-mvp/entities"

type CreateOrderRequest struct {
	ProductID int `json:"product_id" validate:"required"`
	Quantity  int `json:"quantity" validate:"required"`
	User      *entities.User
}
