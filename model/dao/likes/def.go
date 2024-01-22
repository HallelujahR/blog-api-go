package likes

import (
	"time"

	"gorm.io/gorm"
)

const TableNameLikes = "likes"

type Like struct {
	ID        int64 `gorm:"column:id;primaryKey;autoIncrement:true;" json:"id"`
	ArticleID int64 `gorm:"column:article_id;not null;index;uniqueIndex:idx_article_user;" json:"article_id"`
	UserID    int64 `gorm:"column:user_id;not null;index;uniqueIndex:idx_article_user;" json:"user_id"`
	CreatedAt int64 `gorm:"column:created_at;autoCreateTime;" json:"created_at"`
}

func (*Like) TableName() string {
	return TableNameLikes
}

func (u *Like) BeforeCreate(tx *gorm.DB) error {
	// 在插入记录之前自动填充创建时间
	u.CreatedAt = time.Now().Unix()
	return nil
}
