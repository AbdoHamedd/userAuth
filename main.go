package main

import (
	"Users/database"
	"Users/models"
	"Users/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	database.ConnectToDataBase()

	models.Migrate()

	routes.Routes(r)

	r.Run(":6060")

}
