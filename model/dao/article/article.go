package dao

import (
	"blog/model/dao"

	"github.com/gin-gonic/gin"
)

func CreateArticle(g *gin.Context, article Article) error {
	db := dao.NewDbDao(g)
	return db.Create(&article).Error
}

func GetArticleByID(g *gin.Context, ID int64) (Article, error) {
	db := dao.NewDbDao(g)
	record := Article{}
	tx := db.Find(&record, ID)
	if tx != nil {
		return record, tx.Error
	}
	return record, nil
}
func GetArticle(g *gin.Context, page, pageSize int) ([]Article, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	db := dao.NewDbDao(g)
	record := []Article{}
	page = (page - 1) * pageSize
	tx := db.Limit(pageSize).Offset(page).Find(&record)
	if tx != nil {
		return record, tx.Error
	}
	return record, nil
}

func UpdateArticleByID(g *gin.Context, ID int64, data map[string]any) (map[string]any, error) {
	db := dao.NewDbDao(g)
	tx := db.Model(&Article{}).Where("id = ?", ID).Updates(data)
	if tx != nil {
		return data, tx.Error
	}
	return data, nil
}
