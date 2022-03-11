package main

import "github.com/gin-gonic/gin"

func main() {
	route := gin.Default()
	route.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "With love from OCI Devops !",
		})
	})
	route.Run(":8080")
}