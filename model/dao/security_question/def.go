package dao

import (
	"time"

	"gorm.io/gorm"
)

const TableNameSq = "security_question"
const TableNameUsq = "user_security_question"

type SecurityQuestion struct {
	ID          int64  `gorm:"column:id;primaryKey;autoIncrement:true;" json:"id"`
	Question    string `gorm:"column:question;type:varchar(256);not null;" json:"question"`
	UpdatedTime int32  `gorm:"column:updated_time;" json:"updated_time"`
	CreatedTime int32  `gorm:"column:created_time;" json:"created_time"`
}

type UserSecurityQuestion struct {
	ID          int64  `gorm:"column:id;primaryKey;autoIncrement:true;" json:"id"`
	UserID      int64  `gorm:"column:user_id;type:bigint(20);not null;" json:"user_id"`
	QuestionID  int64  `gorm:"column:question_id;type:bigint(20);not null;" json:"question"`
	Answer      string `gorm:"column:answer;type:varchar(256);not null;" json:"answer"`
	UpdatedTime int32  `gorm:"column:updated_time;" json:"updated_time"`
	CreatedTime int32  `gorm:"column:created_time;" json:"created_time"`
}

// TableName QuestionMark's table name
func (*SecurityQuestion) TableName() string {
	return TableNameSq
}

func (u *SecurityQuestion) BeforeCreate(tx *gorm.DB) error {
	// 在插入记录之前自动填充创建时间
	u.CreatedTime = int32(time.Now().Unix())
	u.UpdatedTime = int32(time.Now().Unix())
	return nil
}

func (u *SecurityQuestion) BeforeUpdate(tx *gorm.DB) error {
	u.UpdatedTime = int32(time.Now().Unix())
	return nil
}

// TableName QuestionMark's table name
func (*UserSecurityQuestion) TableName() string {
	return TableNameUsq
}

func (usq *UserSecurityQuestion) BeforeCreate(tx *gorm.DB) error {
	// 在插入记录之前自动填充创建时间
	usq.CreatedTime = int32(time.Now().Unix())
	usq.UpdatedTime = int32(time.Now().Unix())
	return nil
}

func (usq *UserSecurityQuestion) BeforeUpdate(tx *gorm.DB) error {
	usq.UpdatedTime = int32(time.Now().Unix())
	return nil
}
