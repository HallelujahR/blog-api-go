package dao

import (
	"blog/model/dao"

	"github.com/gin-gonic/gin"
)

// 获取密保问题
func GetSecurityQuestion(g *gin.Context) (*[]SecurityQuestion, error) {
	db := dao.NewDbDao(g)
	record := []SecurityQuestion{}
	tx := db.Find(&record)
	if tx != nil {
		return &record, tx.Error
	}

	return &record, nil
}

// 保存密保问题
func SaveSecurityQuestion(g *gin.Context, security_question []UserSecurityQuestion) error {
	db := dao.NewDbDao(g)
	tx := db.Create(&security_question)
	if tx != nil {
		return tx.Error
	}
	return nil
}
