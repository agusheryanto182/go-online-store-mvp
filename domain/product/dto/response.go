package dto

import (
	"github.com/agusheryanto182/go-online-store-mvp/entities"
)

type CreateProductResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Category    int    `json:"category_id"`
}

func GetProductByID(product *entities.Product) CreateProductResponse {
	response := CreateProductResponse{}
	response.ID = product.ID
	response.Name = product.Name
	response.Price = product.Price
	response.Description = product.Description
	response.Category = product.CategoryID

	return response
}

func GetProducts(products []*entities.Product) []CreateProductResponse {
	if len(products) == 0 {
		return []CreateProductResponse{}
	}

	var productsFormatter []CreateProductResponse

	for _, product := range products {
		formatter := GetProductByID(product)
		productsFormatter = append(productsFormatter, formatter)
	}
	return productsFormatter
}
