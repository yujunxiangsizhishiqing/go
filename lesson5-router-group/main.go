package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHander(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "get",
	})
}

func PostHander(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "psot",
	})
}

func DeleteHander(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "delete",
	})
}

func main() {
	r := gin.Default()

	//r.GET("/posts", GetHander)
	//r.POST("/posts", PostHander)
	////删除id=1的文章
	//r.DELETE("/posts/1", DeleteHander)

	//p := r.Group("/posts")
	//p.GET("", GetHander)
	//p.POST("", PostHander)
	////删除id=1的文章
	//p.DELETE("/1", DeleteHander)

	////路由分组:嵌套
	api := r.Group("/posts/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/get", GetHander)
			v1.POST("/post", PostHander)
			v1.DELETE("/delete/1", DeleteHander)
		}

		v2 := api.Group("/v2")
		{
			v2.GET("/get", GetHander)
			v2.POST("/post", PostHander)
			v2.DELETE("/delete/2", DeleteHander)
		}
	}
	r.Run()
}
