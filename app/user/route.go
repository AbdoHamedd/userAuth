package user

import "github.com/gin-gonic/gin"

func Routes(router *gin.Engine) {
	r := router.Group("/app")
	r.POST("/signup", signup)
	r.POST("/login", login)
	r.POST("/logout", logout)
}
