package userservice

import (
	"blog/library"
	"blog/library/resource"
	dao "blog/model/dao/user"
	"errors"

	"github.com/gin-gonic/gin"
)

type ResetPasswordInput struct {
	Key             string `json:"key"`
	Phonenumber     string `json:"phonenumber"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
	Token           string `json:"token"`
}

// 重制密码
func ResetPassword(g *gin.Context) (any, error) {
	var input ResetPasswordInput
	if err := g.BindJSON(&input); err != nil {
		return nil, err
	}

	// 验证token
	s := resource.RedisClient.Get(input.Token)
	if s.Val() != "" {
		return nil, errors.New("修改密码已过期请重新填写密保问题")
	}
	//验证两次密码是否一致
	if input.Password != input.PasswordConfirm {
		return nil, errors.New("两次密码不一致")
	}
	//密码加密
	password, err := library.HashString(input.Password)
	if err != nil {
		return nil, err
	}
	//数据库修改密码
	data, err := dao.UpdateUserByPhone(g, input.Phonenumber, map[string]interface{}{"password": password})
	if err != nil {
		return nil, err
	}
	return data, nil
}
