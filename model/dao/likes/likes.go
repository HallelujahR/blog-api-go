package dao

import (
	"blog/model/dao"

	"github.com/gin-gonic/gin"
)

func GetLikeByArticleID(g *gin.Context, articleID int64) (Like, error) {
	db := dao.NewDbDao(g)
	record := Like{}
	tx := db.Where("article_id = ?", articleID).Find(&record)
	if tx != nil {
		return record, tx.Error
	}
	return record, nil
}
