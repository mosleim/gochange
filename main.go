package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"selamat": "siang",
		})
	})
	route.GET("/pagi", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"selamat": "pagi!",
		})
	})
	route.Run()
}
