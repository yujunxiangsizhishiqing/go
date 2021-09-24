package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//func myPingHnadler(c *gin.Context) {
//	c.JSON(http.StatusOK, gin.H{
//		"message": "pong",
//	})
//}

type User struct {
	ID string `uri:"id" binding:"required"`
	//Name string `uri:"username" binding:"required"`
}

func main() {
	r := gin.Default()
	//r.GET("/ping", myPingHnadler)
	r.POST("/user/:id", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindUri(&user); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"Code": 400,
				"err":  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"Code": 0,
			//"Name": user.Name,
			"ID": user.ID,
		})
	})
	r.Run()
}
