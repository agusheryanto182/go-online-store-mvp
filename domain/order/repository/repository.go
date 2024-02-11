package repository

import (
	"github.com/agusheryanto182/go-online-store-mvp/domain/order"
	"github.com/agusheryanto182/go-online-store-mvp/entities"
	"gorm.io/gorm"
)

type OrderRepositoryImpl struct {
	DB *gorm.DB
}

func NewOrderRepository(DB *gorm.DB) order.OrderRepositoryInterface {
	return &OrderRepositoryImpl{DB: DB}
}

func (r *OrderRepositoryImpl) CreateOrder(newOrder *entities.Order) (*entities.Order, error) {
	err := r.DB.Create(newOrder).Error
	if err != nil {
		return nil, err
	}
	return newOrder, nil
}

func (r *OrderRepositoryImpl) UpdateOrder(ID int, updateOrder *entities.Order) (*entities.Order, error) {
	if err := r.DB.Table("orders").Where("id = ?", ID).Updates(updateOrder).Error; err != nil {
		return nil, err
	}

	return updateOrder, nil
}
