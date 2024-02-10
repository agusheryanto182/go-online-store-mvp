package dto

import (
	"github.com/agusheryanto182/go-online-store-mvp/entities"
)

type CreateProductResponse struct {
	ID          int                     `json:"id"`
	Name        string                  `json:"name"`
	Price       int                     `json:"price"`
	Description string                  `json:"description"`
	Address     string                  `json:"address"`
	Category    string                  `json:"category"`
	Image       []ProductImageFormatter `json:"image"`
}

type ProductImageFormatter struct {
	ID  int    `json:"id"`
	URL string `json:"image_url"`
}

func GetProductByID(product *entities.Product) CreateProductResponse {
	response := CreateProductResponse{}
	response.ID = product.ID
	response.Name = product.Name
	response.Price = product.Price
	response.Description = product.Description
	response.Address = product.Address
	response.Category = product.Category.Name

	var productImages []ProductImageFormatter
	for _, productImage := range product.ProductPhotos {
		if productImage.DeletedAt == nil {
			image := ProductImageFormatter{
				ID:  productImage.ID,
				URL: productImage.ImageURL,
			}
			productImages = append(productImages, image)
		}
	}
	response.Image = productImages

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

type CreateImageProductFormatter struct {
	ID  int    `json:"id"`
	URL string `json:"image"`
}

func CreateImageProductResponse(productImage *entities.ProductPhotos) CreateImageProductFormatter {
	response := CreateImageProductFormatter{}
	response.ID = productImage.ID
	response.URL = productImage.ImageURL

	return response
}
