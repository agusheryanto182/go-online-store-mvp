package service

import (
	"errors"
	"strconv"

	"github.com/agusheryanto182/go-online-store-mvp/domain/user"
	"github.com/agusheryanto182/go-online-store-mvp/entities"
)

type UserServiceImpl struct {
	userRepository user.UserRepositoryInterface
}

func NewUserService(userRepository user.UserRepositoryInterface) user.UserServiceInterface {
	return &UserServiceImpl{userRepository: userRepository}
}

func (s *UserServiceImpl) GetID(ID int) (*entities.User, error) {
	result, err := s.userRepository.FindID(ID)
	if err != nil {
		return nil, errors.New("id with " + strconv.Itoa(ID) + "is not found")
	}
	return result, nil
}

func (s *UserServiceImpl) GetEmail(email string) (*entities.User, error) {
	result, err := s.userRepository.FindEmail(email)
	if err != nil {
		return nil, errors.New("email with " + email + "is not found")
	}
	return result, nil
}

func (s *UserServiceImpl) GetUsername(username string) (*entities.User, error) {
	result, err := s.userRepository.FindUsername(username)
	if err != nil {
		return nil, errors.New("username with " + username + "is not found")
	}
	return result, nil
}
