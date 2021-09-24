package main

import (
	v1 "lesson39-login/api/v1"
	"lesson39-login/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//初始化数据库
	if err := model.InitializeDatabase(); err != nil {
		panic(err)
	}

	r := gin.Default()
	v := r.Group("api/v1")
	{
		//ping
		v.GET("/ping", func(c *gin.Context) {
			c.JSONP(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		//注册
		v.POST("user/register", v1.UserRegisetrHandler)
		//登录
		v.POST("user/login", v1.UserLoginHandler)

	}

	r.Run()
}
