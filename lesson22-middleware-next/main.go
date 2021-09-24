package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func middleware1(c *gin.Context) {
	log.Println("middleware1 in ...")
	c.Set("key", 1000)
	log.Println("before next...")
	c.Next()
	log.Println("next after...")
	log.Println("middleware done...")
}

func main() {
	r := gin.Default()
	r.Use(middleware1)
	r.GET("/ping", func(c *gin.Context) {
		log.Println("func in....")
		k := c.GetInt("key")
		c.Set("key", k+2000)
		log.Println("func done....")
		c.JSON(http.StatusOK, gin.H{
			"message": "213",
			"key":     k,
		})
	})
	r.Run()
}
