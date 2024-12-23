package user

import (
	"Users/response"
	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	err, req := signupValidate(c)
	if !err {
		return
	}
	user := signupServices(req)

	response.Created(c, user)
}

func login(c *gin.Context) {
	user, check := loginvalidation(c)
	if !check {
		return
	}
	loginServices(&user)
	response.Ok(c, user)
}

func logout(c *gin.Context) {
	user := logoutValidate(c)
	if user == nil {
		return
	}
	logoutServices(user)
	response.Ok(c, "Logged Out")
}
