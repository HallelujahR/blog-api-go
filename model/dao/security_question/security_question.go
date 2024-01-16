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
func SaveSecurityQuestion(g *gin.Context, security_question []*UserSecurityQuestion) error {
	db := dao.NewDbDao(g)
	tx := db.Create(&security_question)
	if tx != nil {
		return tx.Error
	}
	return nil
}

// 获取指定用户的密保问题
func GetUserSecurityQuestion(g *gin.Context, userID int64) (*[]UserSecurityQuestion, error) {
	db := dao.NewDbDao(g)
	record := []UserSecurityQuestion{}
	tx := db.Where("user_id = ?", userID).Find(&record)
	if tx != nil {
		return &record, tx.Error
	}
	return &record, nil
}

// 根据ID获取密保问题
func GetSecurityQuestionByIDs(g *gin.Context, id []int64) (*[]SecurityQuestion, error) {
	db := dao.NewDbDao(g)
	record := []SecurityQuestion{}
	tx := db.Find(&record, id).Error
	if tx != nil {
		return &record, tx
	}
	return &record, nil
}
