package main

import (
	fume "github.com/fumeapp/gin"
	"github.com/gin-gonic/gin"
)

func main() {
	routes := gin.New()
	routes.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Gin running with Fume"})
	})
	fume.Start(routes, fume.Options{})
}