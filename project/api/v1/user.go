package v1

import (
	"net/http"
	"project/service"

	"github.com/gin-gonic/gin"
)

// UserLoginHandler 登录
func UserLoginHandler(c *gin.Context) {
	var s service.UserLoginService
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		//return
	} else {
		res := s.Login()
		c.JSON(http.StatusOK, res)
	}

}

// UserRegisterHandler 注册
func UserRegisterHandler(c *gin.Context) {
	var s service.UserRegisterService
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(200, gin.H{
			"msg": "err",
		})
	} else {
		res := s.Register()
		c.JSON(200, res)
	}
}
