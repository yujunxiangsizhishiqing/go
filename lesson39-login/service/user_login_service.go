package service

import (
	"lesson39-login/model"
	"lesson39-login/serializer"
)

type UserLoginService struct {
	Email    string `json:"email" binding:"required,email"`
	PassWord string `json:"password" binding:"required,len=8"`
}

func (service *UserLoginService) Login() serializer.Response {

	sqlStr := "select count(id) from user where email = ? and password = ?"
	var count int
	_ = model.DB.Get(&count, sqlStr, service.Email, service.PassWord)

	if count == 0 {
		return serializer.Response{
			Code: 40003,
			Msg:  "账号或密码错误",
		}
	}
	return serializer.Response{
		Code: 0,
		Msg:  "登录成功",
	}
}
