package repository

import (
	"time"

	"github.com/agusheryanto182/go-online-store-mvp/domain/product"
	"github.com/agusheryanto182/go-online-store-mvp/entities"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) product.ProductRepositoryInterface {
	return &ProductRepositoryImpl{DB: DB}
}

func (r *ProductRepositoryImpl) InsertProduct(productData *entities.Product) (*entities.Product, error) {
	if err := r.DB.Create(&productData).Error; err != nil {
		return nil, err
	}
	return productData, nil
}

func (r *ProductRepositoryImpl) UpdateProduct(ID int, updateProduct *entities.Product) (*entities.Product, error) {
	var product *entities.Product
	if err := r.DB.Where("ID = ? AND deleted_at IS NULL", ID).First(&product).Error; err != nil {
		return nil, err
	}

	if err := r.DB.Model(&product).Updates(updateProduct).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (r *ProductRepositoryImpl) DeleteProduct(ID int) error {
	product := &entities.Product{}
	if err := r.DB.First(product, ID).Error; err != nil {
		return err
	}

	if err := r.DB.Model(product).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}

	return nil
}

func (r *ProductRepositoryImpl) FindAllProducts() ([]*entities.Product, error) {
	var products []*entities.Product
	if err := r.DB.Table("products").Where("deleted_at IS NULL").First(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepositoryImpl) FindProductByID(ID int) (*entities.Product, error) {
	var product *entities.Product
	if err := r.DB.Table("products").Where("ID = ? AND deleted_at IS NULL", ID).First(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r *ProductRepositoryImpl) FindProductByCategoryID(categoryID int) ([]*entities.Product, error) {
	var products []*entities.Product
	if err := r.DB.Where("category_id = ? AND deleted_at IS NULL", categoryID).Preload("Category").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
