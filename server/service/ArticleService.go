package service

import (
	"blog/server/dao"
	"blog/server/model"
)

func GetArticleById(Id int) (article model.Article) {
	if Id == 0 {
		panic(model.GetException(model.ErrorId))
	}
	return dao.GetArticleById(Id)
}
func PageArticleByCategoryPrefix(prefix string, page *model.Page) {

	category := dao.GetCategoryByPrefix(prefix)

	if page.TotalCount = dao.CountArticleByCategoryId(category.Id, model.ArticleStatusPublish); page.TotalCount == 0 {
		return
	}
	page.Data = dao.PageArticleByCategoryId(category.Id, page.PageNo, page.PageSize, model.ArticleStatusPublish)
}
