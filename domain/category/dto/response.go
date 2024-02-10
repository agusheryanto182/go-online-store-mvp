package dto

import "github.com/agusheryanto182/go-online-store-mvp/entities"

type CreateCategoryResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func CreateCategoryResponseFormatter(category *entities.Category) CreateCategoryResponse {
	response := CreateCategoryResponse{}
	response.Id = category.ID
	response.Name = category.Name
	response.Description = category.Description
	response.Image = category.Image
	return response
}

func GetCategoryResponseFormatter(category *entities.Category) CreateCategoryResponse {
	response := CreateCategoryResponse{}
	response.Id = category.ID
	response.Name = category.Name
	response.Description = category.Description
	response.Image = category.Image
	return response
}
