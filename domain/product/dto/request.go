package dto

type CreateProductRequest struct {
	Name        string `json:"name" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Description string `json:"description" validate:"required"`
	CategoryID  int    `json:"category_id" validate:"required"`
}

type UpdateProductRequest struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	CategoryID  int    `json:"category_id"`
}
