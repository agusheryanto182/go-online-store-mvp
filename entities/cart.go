package entities

type Cart struct {
	Id         int         `gorm:"column:id;primaryKey" json:"id"`
	UserId     int         `gorm:"column:user_id" json:"user_id"`
	GrandTotal int         `gorm:"column:grand_total" json:"grand_total"`
	User       *User       `gorm:"foreignKey:UserId" json:"user"`
	CartItems  []*CartItem `gorm:"foreignKey:CartId" json:"cart_items,omitempty"`
}

type CartItem struct {
	Id         int      `gorm:"column:id;primaryKey" json:"id"`
	CartId     int      `gorm:"column:cart_id" json:"cart_id"`
	ProductId  int      `gorm:"column:product_id" json:"product_id"`
	Quantity   int      `gorm:"column:quantity" json:"quantity"`
	Price      int      `gorm:"column:price" json:"price"`
	TotalPrice int      `gorm:"column:total_price" json:"total_price"`
	Product    *Product `gorm:"foreignKey:ProductId" json:"product,omitempty"`
}
