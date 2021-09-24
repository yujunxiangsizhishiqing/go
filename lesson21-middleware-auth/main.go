package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func loginAuth(c *gin.Context) {
	fmt.Println("我是登录保护中间件：loginAuth")
}

func main() {
	r := gin.New()

	////全局调用中间件
	//r.Use(loginAuth)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "token",
		})
	})

	r.POST("/register", func(c *gin.Context) {

	})

	//登录保护
	user := r.Group("/user", loginAuth)
	{
		user.GET("/:id", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "获取用户详情接口，需要登录保护",
			})
		})
		user.PUT("/:id", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "更新用户详情接口，需要登录保护",
			})
		})
	}

	r.Run()
}
