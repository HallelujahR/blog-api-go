package usersecurityquestionservice

import (
	dao "blog/model/dao/security_question"

	"github.com/gin-gonic/gin"
)

func GetSecurityQuestion(g *gin.Context) (any, error) {
	data, err := dao.GetSecurityQuestion(g)
	if err != nil {
		return nil, err
	}
	return data, nil

}

// 处理密保问题
func HandleSecurityQuestion(g *gin.Context, userID int64, securityQuestion []dao.UserSecurityQuestion) error {
	//添加userid
	for _, v := range securityQuestion {
		v.UserID = userID
	}
	err := dao.SaveSecurityQuestion(g, securityQuestion)
	if err != nil {
		return err
	}
	return nil
}
