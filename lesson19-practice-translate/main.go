package main

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	_ "github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	_ "github.com/go-playground/universal-translator"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

type Login struct {
	UserName   string `json:"UserName" binding:"required,min=6,max=12"`
	PassWord   string `json:"PassWord" binding:"required,min=6,max=20"`
	RePassWord string `json:"Re-PassWord" binding:"required,ming=6,max=20"`
}

type Register struct {
	UserName   string `json:"UserName" binding:"required,min=3,max=12"`
	PassWord   string `json:"PassWord" binding:"required,min=3,max=20"`
	RePassWord string `json:"Re-PassWord" binding:"required,min=3,max=20"`
	Age        uint32 `json:"Age" binding:"required,gte=1,lte=120"`
	Sex        uint32 `json:"Sex" binding:"required"`
	Email      string `json:"Email" binding:"required,email",`

}

var trans ut.Translator

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

		err,ok :=err.(validator.ValidationErrors)
		if !ok{
			c.JSON(http.StatusOK, gin.H{
				"Code":    404,
				"Message": "register fail",
				"Err":     err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"code": 40004,
			"err":err.Translate(trans),
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

func initializeTrans()(err error){
	//Accept-Language
	// 修改gin框架validator引擎属性
	if v,ok := binding.Validator.Engine().(*validator.Validate);ok{
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := field.Tag.Get("json")
			return name
		})
		zh := zh.New()
		uni := ut.New(zh,zh)
		trans,_ = uni.GetTranslator("zh")
		err = zhTranslations.RegisterDefaultTranslations(v, trans)
		return
	}
	return
}

func main() {
	if err := initializeTrans();err !=nil{
		fmt.Println(err.Error())
		panic(err)
	}

	r := gin.Default()

	r.POST("/userlogin", loginHandler)
	r.POST("/register", registerHandler)
	r.Run()
}

