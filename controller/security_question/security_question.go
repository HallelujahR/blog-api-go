package securityquestion

import (
	usersecurityquestionservice "blog/model/service/user_security_question"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSecurityQuestion(g *gin.Context) {
	data, err := usersecurityquestionservice.GetSecurityQuestion(g)
	if err != nil {
		returnData := map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
			"data":    data,
		}
		g.JSON(http.StatusInternalServerError, returnData)
		return
	}
	returnData := map[string]interface{}{
		"message": "successs",
		"status":  http.StatusOK,
		"data":    data,
	}
	g.JSON(http.StatusOK, returnData)
}

// 获取用户对应密保问题
func GetSecurityQuestionByUser(g *gin.Context) {
	data, err := usersecurityquestionservice.GetSecurityQuestionByUser(g)
	if err != nil {
		returnData := map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
			"data":    data,
		}
		g.JSON(http.StatusInternalServerError, returnData)
		return
	}
	returnData := map[string]interface{}{
		"message": "successs",
		"status":  http.StatusOK,
		"data":    data,
	}
	g.JSON(http.StatusOK, returnData)
}

// 检查密保问题是否正确
func CheckSecurityQuestion(g *gin.Context) {
	data, err := usersecurityquestionservice.CheckSecurityQuestion(g)
	if err != nil {
		returnData := map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
			"data":    data,
		}
		g.JSON(http.StatusInternalServerError, returnData)
		return
	}
	returnData := map[string]interface{}{
		"message": "successs",
		"status":  http.StatusOK,
		"data":    data,
	}
	g.JSON(http.StatusOK, returnData)
}
