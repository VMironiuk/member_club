package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", rootRequest)
	r.Run()
}

func rootRequest(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
