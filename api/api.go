package api

import "github.com/gin-gonic/gin"

func ruchka() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})
	r.Run(":8080")
}
