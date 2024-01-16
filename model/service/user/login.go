package userservice

import (
	"blog/library/jwt"
	dao "blog/model/dao/user"
	"errors"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginInputParams struct {
	Phonenumber string `json:"phonenumber"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type LoginOutputParams struct {
	dao.User
	Token string `json:"token"`
}

// 用户登录
func Login(g *gin.Context) (any, error) {

	var input LoginInputParams
	if err := g.BindJSON(&input); err != nil {
		return nil, err
	}

	//获取数据
	userData, err := dao.GetUserByCond(g, input.Phonenumber, input.Email)
	if err != nil {
		return nil, err
	}
	//验证密码
	if err := CheckPasswordHash(input.Password, userData.Password); err != nil {
		return nil, err
	}
	//密码验证通过后生成token 使用phonenumber生成
	token, err := jwt.GenerateToken(userData.Phonenumber)

	if err != nil {
		return nil, err
	}

	LoginOutputParams := LoginOutputParams{
		User:  *userData,
		Token: token,
	}

	return LoginOutputParams, nil

}

// 验证密码
func CheckPasswordHash(password string, passwordHash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		return errors.New("密码错误！")
	}
	return nil
}
