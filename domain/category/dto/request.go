package dto

type CreateCategoryRequest struct {
	Name        string `form:"name" validate:"required"`
	Description string `form:"description" validate:"required"`
	Image       string `form:"image"`
}

type UpdateCategoryRequest struct {
	Name        string `form:"name"`
	Description string `form:"description"`
	Image       string `form:"image"`
}
