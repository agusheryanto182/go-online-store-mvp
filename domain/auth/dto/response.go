package dto

import "github.com/agusheryanto182/go-online-store-mvp/entities"

type Login struct {
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Token    string `json:"access_token"`
}

func LoginResponse(user *entities.User, token string) *Login {
	userFormatter := &Login{}
	userFormatter.Username = user.Username
	userFormatter.Avatar = user.Avatar
	userFormatter.Email = user.Email
	userFormatter.Role = user.Role
	userFormatter.Token = token

	return userFormatter
}
