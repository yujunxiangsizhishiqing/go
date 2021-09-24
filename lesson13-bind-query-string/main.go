package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       string `form:"id" binding:"required"`
	Name     string `form:"username" binding:"required"`
	PassWord string `form:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindQuery(&user); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"Code":    400,
				"Message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"Code":     0,
			"ID":       user.ID,
			"UserName": user.Name,
			"PassWord": user.PassWord,
		})
	})
	r.Run()
}
