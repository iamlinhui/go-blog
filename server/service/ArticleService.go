package service

import (
	"blog/server/dao"
	"blog/server/model"
)

func GetArticleById(Id *int) (article model.Article) {
	if *Id == 0 {
		panic(model.GetException(model.ErrorId))
	}
	return dao.GetArticleById(Id)
}
