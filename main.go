package main

import (
	"github.com/CalvinAX/GoShop/controllers"
	"github.com/CalvinAX/GoShop/initializers"
	"github.com/CalvinAX/GoShop/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	api := r.Group("/api")

	auth := api.Group("/auth")
	product := api.Group("/product")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// User routes

	auth.POST("/signup", controllers.SignUp)
	auth.POST("/login", controllers.Login)
	auth.GET("/validate", middleware.RequireAuth, controllers.Validate)

	product.POST("/create", middleware.RequireAuth, controllers.CreateProduct)

	r.Run()
}
