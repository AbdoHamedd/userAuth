package routes

import (
	"Users/app/user"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	user.Routes(r)

}
