package main

import (
	"profile/profile/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Registration route
	// router.POST("/register", func(c *gin.Context) {
	// 	// Handle registration logic here
	// 	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
	// })

	// // Login route
	// router.POST("/login", func(c *gin.Context) {
	// 	// Handle login logic here
	// 	c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully"})
	// })

	// Profile route
	// router.GET("/api/profile", func(c *gin.Context) {
	// 	// Handle profile logic here
	// 	c.JSON(http.StatusOK, gin.H{"message": "User profile retrieved "})
	// })
	handler.NewHandler(&handler.Config{
		R: router,
	})

	router.Run(":8080")
}
