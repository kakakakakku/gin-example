package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/users/:name", func(c *gin.Context) {
		name := c.Params.ByName("name")
		c.JSON(200, gin.H{
			"name":    name,
			"message": "hello!",
		})
	})

	r.Run()
}
