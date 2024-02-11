package repository

import (
	"time"

	"github.com/agusheryanto182/go-online-store-mvp/domain/category"
	"github.com/agusheryanto182/go-online-store-mvp/entities"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	DB *gorm.DB
}

func NewCategoryRepository(DB *gorm.DB) category.CategoryRepositoryInterface {
	return &CategoryRepositoryImpl{DB: DB}
}

func (r *CategoryRepositoryImpl) InsertCategory(category *entities.Category) (*entities.Category, error) {
	if err := r.DB.Create(&category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (r *CategoryRepositoryImpl) UpdateCategory(ID int, updateCategory *entities.Category) (*entities.Category, error) {
	var category *entities.Category
	if err := r.DB.Where("id = ? AND deleted_at is NULL", ID).First(&category).Error; err != nil {
		return nil, err
	}

	if err := r.DB.Model(&category).Updates(updateCategory).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (r *CategoryRepositoryImpl) DeleteCategory(ID int) error {
	category := &entities.Category{}
	if err := r.DB.Where("id = ? AND deleted_at is NULL", ID).First(&category).Error; err != nil {
		return err
	}

	if err := r.DB.Model(&category).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}
	return nil
}

func (r *CategoryRepositoryImpl) FindAllCategory() ([]*entities.Category, error) {
	categories := []*entities.Category{}
	if err := r.DB.Where("deleted_at is NULL").Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *CategoryRepositoryImpl) FindCategoryByID(ID int) (*entities.Category, error) {
	category := &entities.Category{}
	if err := r.DB.Where("id = ? AND deleted_at is NULL", ID).First(&category).Error; err != nil {
		return nil, err
	}
	return category, nil
}
