package main

import (
	"SP2/models"
	"SP2/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	models.InitDB()

	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
