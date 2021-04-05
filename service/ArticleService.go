package service

import (
	"go-blog/dao"
	"go-blog/model"
)

func GetArticleById(Id *int) (article model.Article) {
	return dao.GetArticleById(Id)
}
