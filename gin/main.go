package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()
	server.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	server.GET("/users/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
	server.GET("/views/*.html", func(c *gin.Context) {
		page := c.Param(".html")
		c.String(http.StatusOK, page)
	})
	server.Run(":8080")
}
