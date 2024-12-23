package user

import (
	"Users/exchanging/user"
	"Users/models"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func signupServices(req user.SignupRequest) models.User {
	req.Password, _ = hashPassword(req.Password)
	user := setUser(req)
	createModelUser(&user)
	return user
}

func loginServices(user *models.User) {
	user.Token = buildTokenn(user.ID)
	updateUser(user)
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", fmt.Errorf("can't hash this password: %w", err)
	}
	return string(hashedPassword), nil
}

func validatePassword(hashedPassword string, enteredPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(enteredPassword))
}

func logoutServices(user *models.User) {
	logoutHelp(user)
}
