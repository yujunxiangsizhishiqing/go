package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	ID string `form:"id" binding:"required"`
	Name string `form:"username" binding:"required",min=3,max=12`
	PassWord string `form:"password" binding:"required,min=6,max=20"`
}


}
func main() {
	r := gin.Default()
	r.POST("/user",func(c *gin.Context){
		var user User
		if err := c.ShouldBind(&user);err!=nil{
			c.JSON(http.StatusOK,gin.H{
				"Code":400,
				"Message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"Code": 0,
			"ID": user.ID,
			"UserName": user.Name,
			"PassWord": user.PassWord,
		})
	})
	r.Run()
}
