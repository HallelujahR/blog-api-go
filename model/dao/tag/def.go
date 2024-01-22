package dao

import (
	"time"

	"gorm.io/gorm"
)

const TableNameTags = "tags"

type Tags struct {
	ID        int64  `gorm:"column:id;primaryKey;autoIncrement:true;" json:"id"`
	Image     string `gorm:"column:image;" json:"image"`
	Name      string `gorm:"column:name;not null;" json:"name"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime;" json:"created_at"`
	UpdatedAt int64  `gorm:"column:updated_at;autoCreateTime;" json:"updated_at"`
}

func (*Tags) TableName() string {
	return TableNameTags
}

func (u *Tags) BeforeCreate(tx *gorm.DB) error {
	// 在插入记录之前自动填充创建时间
	u.CreatedAt = time.Now().Unix()
	return nil
}
