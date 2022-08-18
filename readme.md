# Gin+Fume example
> Example setup of a Gin project running in Fume

<p align="center">
  <img src="https://github.com/fumeapp/gin-example/blob/production/gin-example.png?raw=true" />
</p>

```go
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
```