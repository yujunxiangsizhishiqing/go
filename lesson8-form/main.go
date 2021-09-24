package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func postHandler(c *gin.Context) {
	message := c.PostForm("message")
	name := c.DefaultPostForm("name", "nil")
	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"name":    name,
	})
}

func mapHandler(c *gin.Context) {
	user := c.PostFormMap("user")
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func arrayHandler(c *gin.Context) {
	ids := c.PostFormArray("ids")
	c.JSON(http.StatusOK, gin.H{
		"ids": ids,
	})
}

func main() {
	r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, map[string]string{
	//		"message": "pong",
	//	})
	//})

	r.POST("/form-post", postHandler)
	r.POST("/form-array", arrayHandler)
	r.POST("/form-map", mapHandler)

	r.Run()
}
