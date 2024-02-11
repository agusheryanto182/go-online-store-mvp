package dto

import "github.com/agusheryanto182/go-online-store-mvp/entities"

type OrderProductResponse struct {
	Id                 int    `gorm:"column:id;primaryKey" json:"id"`
	IdOrder            string `gorm:"column:id_order;type:VARCHAR(255)" json:"id_order"`
	UserId             int    `gorm:"column:user_id" json:"user_id"`
	GrandTotalQuantity int    `gorm:"column:grand_total_quantity" json:"grand_total_quantity"`
	TotalAmountPaid    int    `gorm:"column:total_amount_paid" json:"total_amount_paid"`
	OrderStatus        string `gorm:"column:order_status;type:VARCHAR(255)" json:"order_status"`
	PaymentURL         string `gorm:"column:payment_url;type:VARCHAR(255)" json:"payment_url"`
}

func FormatOrderResponse(order *entities.Order) *OrderProductResponse {
	orderResponse := &OrderProductResponse{
		Id:                 order.Id,
		IdOrder:            order.IdOrder,
		UserId:             order.UserId,
		GrandTotalQuantity: order.GrandTotalQuantity,
		TotalAmountPaid:    order.GrandTotalQuantity,
		OrderStatus:        order.OrderStatus,
		PaymentURL:         order.PaymentURL,
	}
	return orderResponse
}
