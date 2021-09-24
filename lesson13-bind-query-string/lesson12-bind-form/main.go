package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       string `form:"id" binding:"required"`
	UserName string `form:"username" binding:"required,min=3"`
	PassWord string `form:"password" binding:"required,min=6,max=12"`
}

func main() {
	r := gin.Default()
	r.POST("/user", func(c *gin.Context) {
		var user User
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(200, gin.H{
				"Code":    200,
				"Message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"Code":     0,
			"ID":       user.ID,
			"UserName": user.UserName,
			"PassWord": user.PassWord,
		})
	})
	r.Run()
}
