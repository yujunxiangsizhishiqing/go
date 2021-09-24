package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WelcomeHander(context *gin.Context) {
	//first name , last name
	firstName := context.DefaultQuery("firstname", "nil")
	//lastName := context.DefaultQuery("lastname", "nil")
	lastName := context.Query("lastname")
	context.JSON(http.StatusOK, gin.H{
		"fName": firstName,
		"lName": lastName,
	})
}

func ArrayHander(c *gin.Context) {
	ids := c.QueryArray("ids")
	c.JSON(http.StatusOK, gin.H{
		"ids": ids,
	})
}

func MapHnader(c *gin.Context) {
	key := c.QueryMap("user")
	c.JSON(http.StatusOK, gin.H{
		"data": key,
	})
}

func main() {
	r := gin.Default()
	//r.GET("/wangxu", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "pong",
	//	})
	//})

	r.GET("/welcome", WelcomeHander)
	r.GET("/array", ArrayHander)
	r.GET("/map", MapHnader)
	r.Run()
}
