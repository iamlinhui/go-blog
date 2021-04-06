package service

import (
	"blog/src/dao"
	"blog/src/model"
)

func GetArticleById(Id *int) (article model.Article) {
	if *Id == 0 {
		panic(model.GetException(model.ErrorId))
	}
	return dao.GetArticleById(Id)
}
