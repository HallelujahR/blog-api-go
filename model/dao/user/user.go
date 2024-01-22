package dao

import (
	"blog/model/dao"

	"github.com/gin-gonic/gin"
)

// 根据ID获取users表用户数据
func GetUserByID(g *gin.Context, ID int64) (*User, error) {
	db := dao.NewDbDao(g)
	record := User{}
	tx := db.Find(&record, ID)
	if tx != nil {
		return &record, tx.Error
	}

	return &record, nil
}

// 根据phonenumber 、email 获取用户数据
func GetUserByCond(g *gin.Context, phonenumber string, email string) (*User, error) {
	db := dao.NewDbDao(g)
	record := User{}
	tx := db.Where("phonenumber = ?", phonenumber).Or("email = ?", email).Find(&record)
	if tx != nil {
		return &record, tx.Error
	}
	return &record, nil
}

// 创建user数据
func CreateUser(g *gin.Context, user User) (*User, error) {
	db := dao.NewDbDao(g)
	tx := db.Create(&user)
	if tx != nil {
		return &user, tx.Error
	}

	return &user, nil
}

// 根据电话号修改信息
func UpdateUserByPhone(g *gin.Context, phone string, data map[string]any) (map[string]any, error) {
	db := dao.NewDbDao(g)
	tx := db.Model(&User{}).Where("phonenumber = ?", phone).Updates(data)
	if tx != nil {
		return data, tx.Error
	}
	return data, nil
}
