package user

import "github.com/gin-gonic/gin"

func Routes(router *gin.Engine) {
	r := router.Group("/users")
	r.POST("/signup", signup)
	r.POST("/login", login)
	r.POST("/logout", logout)
}
