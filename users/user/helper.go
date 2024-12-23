package user

import (
	"Users/database"
	"Users/exchanging/user"
	"Users/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func setUser(req user.SignupRequest) models.User {
	return models.User{
		Name:     req.UserName,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	}
}

func buildTokenn(id uint) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, _ := token.SignedString([]byte("1a2b3d4o5h6a7m8e6d"))
	return tokenString
}

func logoutHelp(user *models.User) {
	user.Token = ""
	database.DB.Save(&user)
}
