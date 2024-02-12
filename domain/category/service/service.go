package service

import (
	"errors"
	"strconv"

	"github.com/agusheryanto182/go-online-store-mvp/domain/category"
	"github.com/agusheryanto182/go-online-store-mvp/domain/category/dto"
	"github.com/agusheryanto182/go-online-store-mvp/entities"
)

type CategoryServiceImpl struct {
	categoryRepository category.CategoryRepositoryInterface
}

func NewCategoryService(categoryRepository category.CategoryRepositoryInterface) category.CategoryServiceInterface {
	return &CategoryServiceImpl{categoryRepository: categoryRepository}
}

func (s *CategoryServiceImpl) CreateCategory(request *dto.CreateCategoryRequest) (*entities.Category, error) {
	inputData := &entities.Category{
		Name:        request.Name,
		Image:       request.Image,
		Description: request.Description,
	}

	newCategory, err := s.categoryRepository.InsertCategory(inputData)
	if err != nil {
		return nil, errors.New("failed to create category")
	}
	return newCategory, nil
}

func (s *CategoryServiceImpl) UpdateCategory(ID int, request *dto.UpdateCategoryRequest) (*entities.Category, error) {
	category, err := s.categoryRepository.FindCategoryByID(ID)
	if err != nil {
		return nil, errors.New("category with ID " + strconv.Itoa(ID) + " not found")
	}

	if request.Name != "" {
		category.Name = request.Name
	}
	if request.Description != "" {
		category.Description = request.Description
	}
	if request.Image != "" {
		category.Image = request.Image
	}

	updatedCategory, err := s.categoryRepository.UpdateCategory(ID, category)
	if err != nil {
		return nil, errors.New("failed to update category")
	}
	return updatedCategory, nil
}

func (s *CategoryServiceImpl) DeleteCategory(ID int) error {
	_, err := s.categoryRepository.FindCategoryByID(ID)
	if err != nil {
		return errors.New("category with ID " + strconv.Itoa(ID) + " is not found")
	}

	if err := s.categoryRepository.DeleteCategory(ID); err != nil {
		return errors.New("failed to delete category")
	}
	return nil
}

func (s *CategoryServiceImpl) GetAllCategory() ([]*entities.Category, error) {
	categories, err := s.categoryRepository.FindAllCategory()
	if err != nil {
		return nil, errors.New("failed to get categories")
	}
	return categories, nil
}

func (s *CategoryServiceImpl) GetCategoryByID(ID int) (*entities.Category, error) {
	category, err := s.categoryRepository.FindCategoryByID(ID)
	if err != nil {
		return nil, errors.New("categoy with ID " + strconv.Itoa(ID) + " is not found")
	}
	return category, nil
}
