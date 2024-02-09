package service

import (
	"errors"

	"github.com/agusheryanto182/go-online-store-mvp/domain/auth"
	"github.com/agusheryanto182/go-online-store-mvp/domain/auth/dto"
	"github.com/agusheryanto182/go-online-store-mvp/domain/user"
	"github.com/agusheryanto182/go-online-store-mvp/entities"
	"github.com/agusheryanto182/go-online-store-mvp/helper/hashing"
	"github.com/agusheryanto182/go-online-store-mvp/helper/jwt"
)

type AuthServiceImpl struct {
	repository  auth.AuthRepositoryInterface
	userService user.UserServiceInterface
	jwt         jwt.IJwt
	hashing     hashing.HashInterface
}

func NewAuthService(repository auth.AuthRepositoryInterface, userService user.UserServiceInterface, jwt jwt.IJwt, hashing hashing.HashInterface) auth.AuthServiceInterface {
	return &AuthServiceImpl{repository: repository, userService: userService, jwt: jwt, hashing: hashing}
}

func (s *AuthServiceImpl) Register(dataRequest *dto.RegisterRequest) (*entities.User, error) {
	isEmailExist, _ := s.userService.GetEmail(dataRequest.Email)
	if isEmailExist != nil {
		return nil, errors.New("email already exists")
	}

	isUsernameExists, _ := s.userService.GetUsername(dataRequest.Username)
	if isUsernameExists != nil {
		return nil, errors.New("username already exists")
	}

	if dataRequest.Password != dataRequest.PasswordConfirm {
		return nil, errors.New("password is not match")
	}

	hashPassword, err := s.hashing.HashPassword(dataRequest.Password)
	if err != nil {
		return nil, err
	}

	newUser := &entities.User{
		Fullname: dataRequest.Fullname,
		Username: dataRequest.Username,
		Email:    dataRequest.Email,
		Password: hashPassword,
		Avatar:   "https://avatars.githubusercontent.com/u/112523637?v=4",
		Role:     "user",
	}

	user, err := s.repository.InsertUser(newUser)
	if err != nil {
		return nil, errors.New("create account is failed")
	}

	return user, nil
}

func (s *AuthServiceImpl) Login(dataRequest *dto.LoginRequest) (*entities.User, string, error) {
	user, err := s.repository.FindUserByEmailOrUsername(dataRequest.Identifier)
	if err != nil {
		return nil, "", errors.New("user is not found")
	}

	ValidPassword := s.hashing.CheckPasswordHash(dataRequest.Password, user.Password)
	if !ValidPassword {
		return nil, "", errors.New("incorrect password")
	}

	accessSecret, err := s.jwt.GenerateJWT(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, "", err
	}

	return user, accessSecret, nil
}
