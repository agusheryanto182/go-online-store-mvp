package product

import (
	"github.com/agusheryanto182/go-online-store-mvp/domain/product/dto"
	"github.com/agusheryanto182/go-online-store-mvp/entities"
	"github.com/gofiber/fiber/v2"
)

type ProductRepositoryInterface interface {
	InsertProduct(productData *entities.Product) (*entities.Product, error)
	UpdateProduct(ID int, productData *entities.Product) (*entities.Product, error)
	DeleteProduct(ID int) error
	FindAllProducts() ([]*entities.Product, error)
	FindProductByID(ID int) (*entities.Product, error)
	FindProductByCategoryID(categoryID int) ([]*entities.Product, error)
}

type ProductServiceInterface interface {
	CreateProduct(request *dto.CreateProductRequest) (*entities.Product, error)
	UpdateProduct(ID int, productData *dto.UpdateProductRequest) (*entities.Product, error)
	DeleteProduct(ID int) error
	GetAllProducts() ([]*entities.Product, error)
	GetProductByID(ID int) (*entities.Product, error)
	GetProductByCategoryID(categoryID int) ([]*entities.Product, error)
}

type ProductHandlerInterface interface {
	CreateProduct(c *fiber.Ctx) error
	UpdateProduct(c *fiber.Ctx) error
	DeleteProduct(c *fiber.Ctx) error
	GetAllProducts(c *fiber.Ctx) error
	GetProductByCategoryID(c *fiber.Ctx) error
}
