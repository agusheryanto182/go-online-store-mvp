package dto

import (
	"github.com/agusheryanto182/go-online-store-mvp/entities"
)

type CreateProductResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Address     string `json:"address"`
	Category    string `json:"category"`
}

func GetProductByID(product *entities.Product) CreateProductResponse {
	response := CreateProductResponse{}
	response.ID = product.ID
	response.Name = product.Name
	response.Price = product.Price
	response.Description = product.Description
	response.Address = product.Address
	response.Category = product.Category.Name

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
