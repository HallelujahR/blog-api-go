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
