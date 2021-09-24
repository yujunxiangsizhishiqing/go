package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, map[string]string{
	//		"message": "pong",
	//	})
	//})

	r.GET("post/:id/:action", func(c *gin.Context) {
		id := c.Param("id")
		action := c.Param("action")
		c.JSON(http.StatusOK, gin.H{
			"id":      id,
			"action":  action,
			"message": "get id",
		})
	})
	r.Run()
}
