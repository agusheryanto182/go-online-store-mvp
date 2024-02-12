package entities

import "time"

type Product struct {
	ID          int        `gorm:"column:id;primaryKey" json:"id"`
	Name        string     `gorm:"column:name" json:"name"`
	Price       int        `gorm:"column:price" json:"price"`
	Description string     `gorm:"column:description" json:"description"`
	CategoryID  int        `gorm:"column:category_id" json:"category_id"`
	Category    Category   `gorm:"foreignKey:CategoryID" json:"category"`
	CreatedAt   time.Time  `gorm:"column:created_at;type:TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at;type:TIMESTAMP" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
}
