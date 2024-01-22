package articleservice

import (
	dao "blog/model/dao/article"

	"github.com/gin-gonic/gin"
)

type ArticleInput struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

func GetArticles(g *gin.Context) ([]dao.Article, error) {
	var input ArticleInput
	if err := g.BindJSON(&input); err != nil {
		return nil, err
	}
	articles, err := dao.GetArticle(g, input.Page, input.PageSize)

	if err != nil {
		return nil, err
	}
	return articles, nil
}
