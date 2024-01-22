package comments

import (
	"time"

	"gorm.io/gorm"
)

const TableNameComments = "comments"

type Comment struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true;" json:"id"`
	ArticleID int64     `gorm:"column:article_id;not null;index;" json:"article_id"`
	UserID    int64     `gorm:"column:user_id;not null;index;" json:"user_id"`
	Content   string    `gorm:"column:content;type:text;not null;" json:"content"`
	ParentID  *int64    `gorm:"column:parent_id;index;" json:"parent_id"`
	CreatedAt int64     `gorm:"column:created_at;autoCreateTime;" json:"created_at"`
	UpdatedAt int64     `gorm:"column:updated_at;autoUpdateTime;" json:"updated_at"`
	Replies   []Comment `gorm:"foreignKey:ParentID;" json:"replies,omitempty"` // 子评论
}

func (*Comment) TableName() string {
	return TableNameComments
}

func (u *Comment) BeforeCreate(tx *gorm.DB) error {
	// 在插入记录之前自动填充创建时间
	u.CreatedAt = time.Now().Unix()
	u.UpdatedAt = time.Now().Unix()
	return nil
}

func (u *Comment) BeforeUpdate(tx *gorm.DB) error {
	u.UpdatedAt = time.Now().Unix()
	return nil
}
