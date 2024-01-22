package dao

import (
	"blog/model/dao"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetArticleTag(g *gin.Context, articleID, tagID int64, flag string) ([]ArticleTag, error) {
	db := dao.NewDbDao(g)
	record := []ArticleTag{}
	var tx *gorm.DB
	if flag == "article" {
		tx = db.Where("article_id = ?", articleID).Find(&record)
	}
	if flag == "tag" {
		tx = db.Where("tag_id = ?", tagID).Find(&record)
	}

	if tx != nil {
		return record, tx.Error
	}
	return record, nil
}
