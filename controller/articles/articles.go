package articles

import (
	articleservice "blog/model/service/article"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetArticles(g *gin.Context) {
	data, err := articleservice.GetArticles(g)
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
