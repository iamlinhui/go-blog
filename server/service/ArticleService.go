package service

import (
	"blog/server/dao"
	"blog/server/model"
	"strings"
)

func GetArticleById(Id int) (article model.Article) {
	if Id == 0 {
		panic(model.GetException(model.ErrorId))
	}
	return dao.GetArticleById(Id)
}
func PagePublishArticleByCategoryPrefix(prefix string, page *model.Page) {
	prefix = strings.TrimSpace(prefix)
	if len(prefix) == 0 {
		panic(model.GetException(model.ErrorPrefix))
	}
	category := dao.GetCategoryByPrefix(prefix)

	if page.TotalCount = dao.CountArticleByCategoryIdAndStatus(category.Id, model.ArticleStatusPublish); page.TotalCount == 0 {
		return
	}
	page.Data = dao.PageArticleByCategoryIdAndStatus(category.Id, page.PageNo, page.PageSize, model.ArticleStatusPublish)
}

func PagePublishArticle(page *model.Page) {
	if page.TotalCount = dao.CountArticleByStatus(model.ArticleStatusPublish); page.TotalCount == 0 {
		return
	}
	page.Data = dao.PageArticleByStatus(page.PageNo, page.PageSize, model.ArticleStatusPublish)
}
