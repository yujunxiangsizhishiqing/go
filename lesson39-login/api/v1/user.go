package v1

import (
	"lesson39-login/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserLoginHandler(c *gin.Context) {
	var s service.UserLoginService
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSONP(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	} else {
		res := s.Login()
		c.JSON(http.StatusOK, res)
	}
}

func UserRegisetrHandler(c *gin.Context) {
	var s service.UserRegisterService
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		//return
	} else {
		res := s.Register()
		c.JSON(http.StatusOK, res)
		//c.JSON(http.StatusOK, gin.H{
		//	"msg": "success",
		//})
	}

}
