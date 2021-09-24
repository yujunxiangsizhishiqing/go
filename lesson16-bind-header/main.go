package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Header struct {
	Rerferer string `header:"Referrer" binding:"required"`
}

func main() {
	r := gin.Default()

	r.GET("/header", func(c *gin.Context) {
		var header Header
		if err := c.ShouldBindHeader(&header); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"Code": 400,
				"Err":  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"Code":    0,
			"Referer": header.Rerferer,
		})
	})
	r.Run()
}
