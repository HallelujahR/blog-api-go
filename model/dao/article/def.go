package article

import (
	"time"

	"gorm.io/gorm"
)

const TableNameArticle = "articles"

type Article struct {
	ID           int64  `gorm:"column:id;primaryKey;autoIncrement:true;" json:"id"`
	UserID       int64  `gorm:"column:user_id;not null;index;" json:"user_id"`
	Title        string `gorm:"column:title;type:varchar(255);not null;" json:"title"`
	Content      string `gorm:"column:content;type:text;not null;" json:"content"`
	ViewCount    int    `gorm:"column:view_count;default:0;" json:"view_count"`
	LikeCount    int    `gorm:"column:like_count;default:0;" json:"like_count"`
	CommentCount int    `gorm:"column:comment_count;default:0;" json:"comment_count"`
	ShareCount   int    `gorm:"column:share_count;default:0;" json:"share_count"`
	Status       int    `gorm:"column:status;not null;default:100" json:"status"`
	CreatedAt    int64  `gorm:"column:created_at;autoCreateTime;" json:"created_at"`
	UpdatedAt    int64  `gorm:"column:updated_at;autoUpdateTime;" json:"updated_at"`
}

func (*Article) TableName() string {
	return TableNameArticle
}

func (u *Article) BeforeCreate(tx *gorm.DB) error {
	// 在插入记录之前自动填充创建时间
	u.CreatedAt = time.Now().Unix()
	u.UpdatedAt = time.Now().Unix()
	return nil
}

func (u *Article) BeforeUpdate(tx *gorm.DB) error {
	u.UpdatedAt = time.Now().Unix()
	return nil
}
