package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "pong",
	//	})
	//})

	r.POST("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		userName := c.PostForm("user-name")
		c.JSON(http.StatusOK, gin.H{
			"id":        id,
			"user-name": userName,
		})
	})
	r.Run()
}
