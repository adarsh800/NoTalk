package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler struct will ultimately hold "services" properties
// contain methods that implement the application's core functionality
type Handler struct {
}

// Config struct will hold all values needed to initialize the handler package

type Config struct {
	R *gin.Engine
}

// NewHandler "factory" function used to initialize the routes
// and will be called from the main package
func NewHandler(c *Config) {
	// Create a handler (which will later have injected services)
	// h := &Handler{} // currently has no properties

	// Create an account group
	g := c.R.Group("/api/profile")

	// Profile route
	g.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "space persons",
		})
	})
}
