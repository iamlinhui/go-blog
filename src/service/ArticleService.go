package service

import (
	"blog/src/dao"
	"blog/src/model"
)

func GetArticleById(Id *int) (article model.Article) {
	return dao.GetArticleById(Id)
}