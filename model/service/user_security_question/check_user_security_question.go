package usersecurityquestionservice

import (
	"blog/library/resource"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

type CheckInputParams struct {
	Questions   []any  `json:"questions"`
	Phonenumber string `json:"phonenumber"`
	Email       string `json:"email"`
}

// 检查密保问题是否正确
func CheckSecurityQuestion(g *gin.Context) (any, error) {
	var params CheckInputParams
	err := g.BindJSON(&params)
	if err != nil {
		return nil, err
	}
	data, err := HandleQuestionData(g, params.Phonenumber, params.Email, true)
	if err != nil {
		return nil, err
	}
	//进行数据比对 判断是否密保问题正确
	err = CheckAnswer(params.Questions, data)
	if err != nil {
		return nil, err
	}
	//答案全部正确返回token
	token, err := GenerateSecurityQuestionToken(g, params.Phonenumber)
	if err != nil {
		return nil, err
	}
	//存储在redis
	key := params.Phonenumber + "_find_password"
	s := resource.RedisClient.Set(key, token, 5*time.Minute)
	if s.Err() != nil {
		return nil, s.Err()
	}
	returnData := map[string]any{
		"token": token,
		"key":   key,
	}
	return returnData, nil

}

// 对比答案
func CheckAnswer(params any, data []UserQuestionOutputParams) error {

	for _, v := range params.([]any) {
		for _, i := range data {

			if int64(v.(map[string]any)["id"].(float64)) == i.ID {

				if v.(map[string]any)["answer"] == i.Answer {
					return nil
				} else {
					return errors.New("密保答案错误")
				}
			}

		}

	}
	return nil
}

// 生成密保修改密码token
func GenerateSecurityQuestionToken(g *gin.Context, phonenumber string) (string, error) {
	// 创建一个足够长的随机字节序列
	randomBytes := make([]byte, 32) // 32字节的随机数
	if _, err := io.ReadFull(rand.Reader, randomBytes); err != nil {
		return "", err
	}

	// 将随机字节序列编码为 Base64 字符串
	token := base64.URLEncoding.EncodeToString(randomBytes)

	// 可以选择添加用户名和时间戳之类的信息以确保唯一性
	token = fmt.Sprintf("%s:%s:%d", phonenumber, token, time.Now().Unix())

	return token, nil

}
