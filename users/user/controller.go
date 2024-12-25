package user

import (
	"Users/response"
	"Users/transform"
	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	err, req := signupValidate(c)
	if !err {
		return
	}
	user := signupServices(req)
	userTransform := transform.UserTransformSignup(&user)
	response.Created(c, userTransform)
}

func login(c *gin.Context) {
	user, check := loginvalidation(c)
	if !check {
		return
	}
	loginServices(&user)
	userTransform := transform.UserTransformLogin(&user)

	response.Ok(c, userTransform)
}

func logout(c *gin.Context) {
	user := logoutValidate(c)
	if user == nil {
		return
	}
	logoutServices(user)
	response.Ok(c, "Logged Out")
}
