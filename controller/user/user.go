package user

import (
	dao "blog/model/dao/user"
	userservice "blog/model/service/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Test(g *gin.Context) {
	fmt.Println("user controller")
	dao.GetUserByID(g, 1)
}

// 注册用户
func Register(g *gin.Context) {
	//用户注册
	data, err := userservice.Register(g)
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

// 用户登陆
func Login(g *gin.Context) {
	data, err := userservice.Login(g)
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
