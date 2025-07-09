package routes

import (
	"SP2/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to Mini Order System!"})
	})
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

}
