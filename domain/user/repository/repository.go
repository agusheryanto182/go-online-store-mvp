package repository

import (
	"github.com/agusheryanto182/go-online-store-mvp/domain/user"
	"github.com/agusheryanto182/go-online-store-mvp/entities"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) user.UserRepositoryInterface {
	return &UserRepositoryImpl{DB}
}

func (r *UserRepositoryImpl) FindID(ID int) (*entities.User, error) {
	var user *entities.User
	if err := r.DB.Where("id = ? AND deleted_at IS NULL", ID).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) FindEmail(email string) (*entities.User, error) {
	var user *entities.User
	if err := r.DB.Table("users").Where("email = ? AND deleted_at IS NULL", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) FindUsername(username string) (*entities.User, error) {
	var user *entities.User
	if err := r.DB.Where("username = ? AND deleted_at IS NULL", username).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
