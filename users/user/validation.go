package user

import (
	"Users/exchanging/user"
	"Users/models"
	"Users/response"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"regexp"
	"strings"
)

func ValidatePassNameEmail(name, pass, email string) error {
	var err error
	// Validate name
	if len(name) < 2 || len(name) > 50 {
		err = errors.New("user name must be between 2 and 50 characters")
		return err
	}

	// Validate password
	if len(pass) < 2 || len(pass) > 15 {
		err = errors.New("password must be between 2 and 15 characters")
		return err
	}

	// Validate email
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(emailRegex, strings.TrimSpace(email))
	if !matched {
		err = errors.New("invalid email address")
		return err
	}

	return nil
}

func signupValidate(c *gin.Context) (bool, user.SignupRequest) {
	var req user.SignupRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		err = errors.New("validation Erorr")
		response.BadRequest(c, err.Error())
		return false, user.SignupRequest{}
	}
	//check, errorUrl := user.StoreValidate(c.Request, &req)
	//if !check {
	//	for key, value := range errorUrl {
	//		fmt.Println(key, value)
	//	}
	//	fmt.Println("zezooooooooooo")
	//	response.BadRequest(c, errorUrl)
	//	return false, user.SignupRequest{}
	//}
	if err := ValidatePassNameEmail(req.UserName, req.Password, req.Email); err != nil {

		response.BadRequest(c, err.Error())
		return false, user.SignupRequest{}
	}
	if check, _ := emailIsExists(req.Email); check {
		err = errors.New("this user is already exist")
		response.BadRequest(c, err.Error())
		return false, user.SignupRequest{}
	}
	return true, req
}

func loginvalidation(c *gin.Context) (models.User, bool) {
	var req user.LoginRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		err = errors.New("validation error")
		response.BadRequest(c, err.Error())
		return models.User{}, false
	}

	found, user := emailIsExists(req.Email)
	if !found {
		err = errors.New("this user does not exist")
		response.BadRequest(c, err.Error())
		return models.User{}, false
	}

	fmt.Println("Entered Password:", req.Password)
	fmt.Println("Stored Hashed Password:", user.Password)

	passwordError := validatePassword(user.Password, req.Password)
	if passwordError != nil {
		passwordError = errors.New("please enter a valid password")
		response.BadRequest(c, passwordError.Error())
		return models.User{}, false
	}

	return user, true
}

func logoutValidate(c *gin.Context) *models.User {
	var req user.LogoutRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		err = errors.New("validation Erorr")
		response.BadRequest(c, err.Error())
		return &models.User{}
	}
	found, user := userIsFound(req.Id)
	if !found {
		err = errors.New("this user is not found")
		response.NotFound(c, err.Error())
		return &models.User{}
	}
	return &user
}
