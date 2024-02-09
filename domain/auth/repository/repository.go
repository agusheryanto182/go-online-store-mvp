package repository

import (
	"github.com/agusheryanto182/go-online-store-mvp/domain/auth"
	"github.com/agusheryanto182/go-online-store-mvp/entities"
	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	DB *gorm.DB
}

func NewAuthRepository(DB *gorm.DB) auth.AuthRepositoryInterface {
	return &AuthRepositoryImpl{DB: DB}
}

func (r *AuthRepositoryImpl) InsertUser(newUser *entities.User) (*entities.User, error) {
	if err := r.DB.Create(&newUser).Error; err != nil {
		return nil, err
	}
	return newUser, nil
}

func (r *AuthRepositoryImpl) FindUserByEmailOrUsername(identifier string) (*entities.User, error) {
	var user *entities.User
	if err := r.DB.Table("users").Where("(email = ? OR username = ?) AND deleted_at is NULL", identifier, identifier).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
