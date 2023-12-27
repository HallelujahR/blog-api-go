package dao

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUser = "users"

type User struct {
	ID          int64  `gorm:"column:id;primaryKey;autoIncrement:true;" json:"id"`
	Name        string `gorm:"column:name;not null" json:"name"`
	Phonenumber string `gorm:"column:phonenumber;" json:"phonenumber"`
	Email       string `gorm:"column:email;" json:"email"`
	Password    string `gorm:"column:password;" json:"password"`
	Status      int32  `gorm:"column:status;" json:"status"`
	UpdatedTime int32  `gorm:"column:updated_time;" json:"updated_time"`
	CreatedTime int32  `gorm:"column:created_time;" json:"created_time"`
}

// TableName QuestionMark's table name
func (*User) TableName() string {
	return TableNameUser
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	// 在插入记录之前自动填充创建时间
	u.CreatedTime = int32(time.Now().Unix())
	u.UpdatedTime = int32(time.Now().Unix())
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	u.UpdatedTime = int32(time.Now().Unix())
	return nil
}
