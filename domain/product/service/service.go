package service

import (
	"errors"

	"github.com/agusheryanto182/go-online-store-mvp/domain/category"
	"github.com/agusheryanto182/go-online-store-mvp/domain/product"
	"github.com/agusheryanto182/go-online-store-mvp/domain/product/dto"
	"github.com/agusheryanto182/go-online-store-mvp/entities"
)

type ProductServiceImpl struct {
	productRepository product.ProductRepositoryInterface
	categoryService   category.CategoryServiceInterface
}

func NewProductService(productRepository product.ProductRepositoryInterface, categoryService category.CategoryServiceInterface) product.ProductServiceInterface {
	return &ProductServiceImpl{productRepository: productRepository, categoryService: categoryService}
}

func (s *ProductServiceImpl) CreateProduct(request *dto.CreateProductRequest) (*entities.Product, error) {
	_, err := s.categoryService.GetCategoryByID(request.CategoryID)
	if err != nil {
		return nil, errors.New("category not found")
	}

	productData := &entities.Product{
		Name:        request.Name,
		Price:       request.Price,
		Description: request.Description,
		Address:     request.Address,
		CategoryID:  request.CategoryID,
	}

	createdProduct, err := s.productRepository.InsertProduct(productData)
	if err != nil {
		return nil, errors.New("failed to create product")
	}

	return createdProduct, nil
}

func (s *ProductServiceImpl) UpdateProduct(ID int, request *dto.UpdateProductRequest) (*entities.Product, error) {
	product, err := s.productRepository.FindProductByID(ID)
	if err != nil {
		return nil, errors.New("product not found")
	}

	if request.Name != "" {
		product.Name = request.Name
	}

	if request.Price != 0 {
		product.Price = request.Price
	}

	if request.Address != "" {
		product.Address = request.Address
	}

	if request.Description != "" {
		product.Description = request.Description
	}

	if request.CategoryID != 0 {
		_, err := s.categoryService.GetCategoryByID(request.CategoryID)
		if err != nil {
			return nil, errors.New("category not found")
		}
		product.CategoryID = request.CategoryID
	}
	return product, nil
}

func (s *ProductServiceImpl) DeleteProduct(ID int) error {
	product, err := s.productRepository.FindProductByID(ID)
	if err != nil {
		return errors.New("product not found")
	}

	if err := s.productRepository.DeleteProduct(product.ID); err != nil {
		return errors.New("failed to delete product")
	}
	return nil
}

func (s *ProductServiceImpl) GetProductByID(ID int) (*entities.Product, error) {
	product, err := s.productRepository.FindProductByID(ID)
	if err != nil {
		return nil, errors.New("failed to get product by ID")
	}

	return product, nil
}

func (s *ProductServiceImpl) GetProductByCategoryID(categoryID int) ([]*entities.Product, error) {
	products, err := s.productRepository.FindProductByCategoryID(categoryID)
	if err != nil {
		return nil, errors.New("failed to get product by category")
	}
	return products, nil
}
