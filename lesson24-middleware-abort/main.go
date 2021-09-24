package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func middleware1(c *gin.Context) {
	log.Println("middleware1 in ...")
	c.Set("key", 1000)
	log.Println("middleware1 before next...")

	if k := c.GetInt("key"); k == 1000 {
		c.Abort()
		return
	}
	c.Next()
	log.Println("middleware1 next after...")
	log.Println("middleware1 done...")
}

func middleware2(c *gin.Context) {
	log.Println("middleware2 in ...")
	log.Println("middleware2 before next...")
	c.Next()
	log.Println("middleware2 next after...")
	log.Println("middleware2 done...")
}

//多个中间件的next的执行顺序：先进后执行
func main() {
	r := gin.Default()
	r.Use(middleware1)
	r.Use(middleware2)
	r.GET("/ping", func(c *gin.Context) {
		log.Println("func in....")
		k := c.GetInt("key")
		c.Set("key", k+2000)
		log.Println("func done....")
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"key":     k,
		})
	})
	r.Run()
}
