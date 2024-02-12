package category

import (
	"github.com/agusheryanto182/go-online-store-mvp/domain/category/dto"
	"github.com/agusheryanto182/go-online-store-mvp/entities"
	"github.com/gofiber/fiber/v2"
)

type CategoryRepositoryInterface interface {
	InsertCategory(category *entities.Category) (*entities.Category, error)
	UpdateCategory(ID int, updateCategory *entities.Category) (*entities.Category, error)
	DeleteCategory(ID int) error
	FindCategoryByID(ID int) (*entities.Category, error)
	FindAllCategory() ([]*entities.Category, error)
}

type CategoryServiceInterface interface {
	CreateCategory(request *dto.CreateCategoryRequest) (*entities.Category, error)
	UpdateCategory(ID int, request *dto.UpdateCategoryRequest) (*entities.Category, error)
	DeleteCategory(ID int) error
	GetCategoryByID(ID int) (*entities.Category, error)
	GetAllCategory() ([]*entities.Category, error)
}

type CategoryHandlerInterface interface {
	CreateCategory(c *fiber.Ctx) error
	UpdateCategory(c *fiber.Ctx) error
	DeleteCategory(c *fiber.Ctx) error
	GetAllCategory(c *fiber.Ctx) error
}
