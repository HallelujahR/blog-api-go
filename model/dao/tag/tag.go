package dao

import (
	"blog/model/dao"

	"github.com/gin-gonic/gin"
)

func CreateTag(g *gin.Context, tag Tags) error {
	db := dao.NewDbDao(g)
	return db.Create(&tag).Error
}

func GetTags(g *gin.Context, articleID int64) (Tags, error) {
	db := dao.NewDbDao(g)
	record := Tags{}
	tx := db.Find(&record)
	if tx != nil {
		return record, tx.Error
	}
	return record, nil
}

func DeleteTagByID(g *gin.Context, id int64) error {
	db := dao.NewDbDao(g)
	return db.Delete(&Tags{ID: id}).Error
}
