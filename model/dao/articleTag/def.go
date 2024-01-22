package dao

const TableNameArticleTag = "article_tag"

type ArticleTag struct {
	ID        int64 `gorm:"column:id;primaryKey;autoIncrement:true;" json:"id"`
	ArticleID int64 `gorm:"column:article_id;not null;index;" json:"article_id"`
	TagID     int64 `gorm:"column:tag_id;not null;" json:"tag_id"`
}

func (*ArticleTag) TableName() string {
	return TableNameArticleTag
}
