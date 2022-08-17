package main

import (
	fume "github.com/fumeapp/gin"
	"github.com/gin-gonic/gin"
)

func main() {
	routes := gin.New()
	routes.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"message": "Hello World"}) })
	fume.Start(routes, &fume.Options{Dev: false, Port: 8080})
}