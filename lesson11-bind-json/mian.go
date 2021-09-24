package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required,min=3,max=12"`
}

func main() {
	r := gin.Default()
	r.POST("/user", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"Code":    200,
				"Message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"Code": 0,
			"ID":   user.ID,
			"Name": user.Name,
		})
	})

	r.Run()
}
