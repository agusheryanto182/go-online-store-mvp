package user

import (
	"github.com/agusheryanto182/go-online-store-mvp/entities"
)

type UserRepositoryInterface interface {
	FindID(ID int) (*entities.User, error)
	FindEmail(email string) (*entities.User, error)
	FindUsername(username string) (*entities.User, error)
}

type UserServiceInterface interface {
	GetID(ID int) (*entities.User, error)
	GetEmail(email string) (*entities.User, error)
	GetUsername(username string) (*entities.User, error)
}
