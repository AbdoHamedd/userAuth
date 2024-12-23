package routes

import (
	"Users/users/user"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	user.Routes(r)

}
