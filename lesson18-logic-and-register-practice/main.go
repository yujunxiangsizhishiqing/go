package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	UserName   string `json:"UserName" binding:"required",min=6,max=12`
	PassWord   string `json:"PassWord" binding:"required",min=6,max=20`
	RePassWord string `json:"Re-PassWord" binding:"required",ming=6,max=20`
}

type Register struct {
	UserName   string `json:"UserName" binding:"required",min=6,max=12`
	PassWord   string `json:"PassWord" binding:"required",min=6,max=20`
	RePassWord string `json:"Re-PassWord" binding:"required",ming=6,max=20`
	Age        uint32 `json:"Age" binding:"required" gte=1,lte=120`
	Sex        uint32 `json:"Sex" binding:"required"`
	Email      string `json:"Email" binding:"required,email",`
}

func loginHandler(c *gin.Context) {
	var login Login
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Code":    404,
			"message": "fail",
			"err":     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Code":     0,
		"Message":  "success",
		"UserName": login.UserName,
	})
}

func registerHandler(c *gin.Context) {
	var register Register
	if err := c.ShouldBindJSON(&register); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Code":    404,
			"Message": "register fail",
			"Err":     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Code":    0,
		"Message": "register succ",
		"data":    register,
		//"Age":     register.Age,
		//"Sex":     register.Sex,
		//"Email":   register.Email,
	})
}
func main() {
	r := gin.Default()

	r.POST("/userlogin", loginHandler)
	r.POST("/register", registerHandler)
	r.Run()
}
