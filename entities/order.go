package entities

import "time"

type Order struct {
	Id                 int        `gorm:"column:id;primaryKey" json:"id"`
	IdOrder            string     `gorm:"column:id_order;type:VARCHAR(255)" json:"id_order"`
	UserId             int        `gorm:"column:user_id" json:"user_id"`
	GrandTotalQuantity int        `gorm:"column:grand_total_quantity" json:"grand_total_quantity"`
	TotalAmountPaid    int        `gorm:"column:total_amount_paid" json:"total_amount_paid"`
	OrderStatus        string     `gorm:"column:order_status;type:VARCHAR(255)" json:"order_status"`
	PaymentURL         string     `gorm:"column:payment_url;type:VARCHAR(255)" json:"payment_url"`
	CreatedAt          time.Time  `gorm:"column:created_at;type:TIMESTAMP" json:"created_at"`
	UpdatedAt          time.Time  `gorm:"column:updated_at;type:TIMESTAMP" json:"updated_at"`
	DeletedAt          *time.Time `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
	User               User       `gorm:"foreignKey:UserId" json:"user"`
}
