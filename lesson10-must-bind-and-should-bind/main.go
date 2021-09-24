package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID string `form:"id" binding:"required",uuid`
}

func main() {
	r := gin.Default()
	/*shoule bind*/
	//r.GET("/user", func(c *gin.Context) {
	//	var user User
	//	if err := c.ShouldBindQuery(&user); err != nil {
	//		c.JSON(http.StatusOK, gin.H{
	//			"Code":    200,
	//			"Message": err.Error(),
	//		})
	//		return
	//	}
	//	c.JSON(http.StatusOK, gin.H{
	//		"Code": 0,
	//		"ID":   user.ID,
	//	})
	//})

	/*must bind*/
	r.GET("/user", func(c *gin.Context) {
		var user User
		if err := c.BindQuery(&user); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"Code":    200,
				"Message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"Code": 0,
			"ID":   user.ID,
		})
	})
	r.Run()
}
