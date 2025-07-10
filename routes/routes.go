package routes

import (
	"SP2/controllers"
	"SP2/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to Mini Order System!"})
	})
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	// Protected group
	auth := r.Group("/")
	auth.Use(middlewares.AuthMiddleware())

	// 之後會放：商品、訂單等路由
	auth.GET("/protected", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Token passed. You are authorized!"})
	})
}
