package routes

import (
	"Users/users/app"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	app.Routes(r)

}
