package usersecurityquestionservice

import (
	dao "blog/model/dao/security_question"
	userDao "blog/model/dao/user"

	"github.com/gin-gonic/gin"
)

// 获取密保问题
type InputParams struct {
	Phonenumber string `json:"phonenumber"`
	Email       string `json:"email"`
}

type UserQuestionOutputParams struct {
	dao.UserSecurityQuestion
	Question string `json:"question"`
}

func GetSecurityQuestionByUser(g *gin.Context) ([]UserQuestionOutputParams, error) {
	var params InputParams
	err := g.BindJSON(&params)
	//获取用户ID
	userData, err := userDao.GetUserByCond(g, params.Phonenumber, params.Email)
	if err != nil {
		return nil, err
	}
	userID := userData.ID
	//获取所有该用户密保问题ID
	data, err := dao.GetUserSecurityQuestion(g, userID)
	if err != nil {
		return nil, err
	}
	questionIDs := make([]int64, 0)
	for _, v := range *data {
		questionIDs = append(questionIDs, v.QuestionID)
	}
	//根据ID获取密保问题内容
	questionData, err := dao.GetSecurityQuestionByIDs(g, questionIDs)
	//拼接数据把问题内容拼接到密保问题数据中 添加到UserQuestionOutputParams 中返回

	returnData := make([]UserQuestionOutputParams, 0)

	for _, v := range *data {
		for _, q := range *questionData {
			if v.QuestionID == q.ID {
				UserQuestionOutputParams := UserQuestionOutputParams{
					UserSecurityQuestion: v,
					Question:             q.Question,
				}
				returnData = append(returnData, UserQuestionOutputParams)
			}
		}
	}

	return returnData, nil
}
