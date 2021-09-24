package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//func middleware1(c *gin.Context) {
//
//}
//
//func middleware2() gin.HandlerFunc {
//	return func(c *gin.Context) {
//
//	}
//}

func RefererMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ref := c.GetHeader("Referer")
		if ref == "" {
			//c.Abort()
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"message": "非法访问",
			})
			return
		}
		c.Next()
	}
}

func main() {
	r := gin.Default()
	////两种使用中间件的写法
	//r.Use(middleware1)   //1
	//r.Use(middleware2()) //2

	r.Use(RefererMiddleware())
	r.Use(func(c *gin.Context) {
		fmt.Println("我是第二个中间件")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
