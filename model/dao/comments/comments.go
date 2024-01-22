package dao

import (
	"blog/model/dao"

	"github.com/gin-gonic/gin"
)

func CreateComments(g *gin.Context, comment Comment) error {
	db := dao.NewDbDao(g)
	return db.Create(&comment).Error
}

func GetCommentsByArticleID(g *gin.Context, articleID int64) (Comment, error) {
	db := dao.NewDbDao(g)
	record := Comment{}
	tx := db.Where("article_id = ?", articleID).Find(&record)
	if tx != nil {
		return record, tx.Error
	}
	return record, nil
}

func DeleteCommentByID(g *gin.Context, id int64) error {
	db := dao.NewDbDao(g)
	return db.Delete(&Comment{ID: id}).Error
}
