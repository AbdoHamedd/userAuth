package transform

import (
	"Users/exchanging/user"
	"Users/models"
)

func UserTransformSignup(User *models.User) user.UserSignupResponse {
	return user.UserSignupResponse{
		Name:  User.Name,
		Email: User.Email,
		ID:    User.ID,
		Role:  User.Role,
	}
}

func UserTransformLogin(User *models.User) user.UserLoginResponse {
	return user.UserLoginResponse{
		Name:  User.Name,
		Email: User.Email,
		ID:    User.ID,
		Role:  User.Role,
		Token: User.Token,
	}
}
