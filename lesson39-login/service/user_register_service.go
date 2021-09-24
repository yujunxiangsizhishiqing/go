package service

import (
	"lesson39-login/model"
	"lesson39-login/serializer"
)

type UserRegisterService struct {
	NickName string `json:"nickname" binding:"required,min=3,max=7"`
	PassWord string `json:"password" binding:"required,len=8"`
	Age      uint32 `json:"age" binding:"required,gte=1,lte=150"`
	Sex      uint32 `json:"sex" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

func (service *UserRegisterService) Register() serializer.Response {
	//验证邮箱是否已经注册
	sqlStr := "select count(id) from user where email = ?"
	var count int
	_ = model.DB.Get(&count, sqlStr, service.Email)
	if count > 0 {
		return serializer.Response{
			Code:  40001,
			Data:  nil,
			Msg:   "邮箱已注册",
			Error: "",
		}
	}

	//密码加密,先不学习

	//创建用户
	sqlStr2 := "insert into user (nickname,password,age,sex,email) values(:nickname,:password,:age,:sex,:email)"
	_, err := model.DB.NamedExec(sqlStr2, service)
	if err != nil {
		return serializer.Response{
			Code:  40002,
			Data:  nil,
			Msg:   "注册失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Code:  0,
		Data:  nil,
		Msg:   "注册成功",
		Error: "",
	}
}
