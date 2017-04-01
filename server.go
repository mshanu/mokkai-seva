package main

import "gopkg.in/gin-gonic/gin.v1"

func main() {
	router := gin.Default()

	router.GET("/users", func(c *gin.Context) {
		c.String(200, "hello world")
	})
	router.Run(":8080")
}
