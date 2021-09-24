package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//r := gin.Default()
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
