package service

import (
	"blog/server/dao"
	"blog/server/model"
	"strings"
)

func GetArticleById(id int) (article model.Article) {
	if id == 0 {
		panic(model.GetException(model.ErrorId))
	}
	return dao.GetArticleById(id)
}

func GetPublishArticleByUrl(url string) (article model.Article) {
	if url = strings.TrimSpace(url); len(url) == 0 {
		panic(model.GetException(model.ErrorUrl))
	}
	return dao.GetArticleByIdAndStatus(url, model.ArticleStatusPublish)
}
func PagePublishArticleByCategoryPrefix(prefix string, page *model.Page) {
	if prefix = strings.TrimSpace(prefix); len(prefix) == 0 {
		panic(model.GetException(model.ErrorPrefix))
	}
	category := dao.GetCategoryByPrefix(prefix)

	if page.TotalCount = dao.CountArticleByCategoryIdAndStatus(category.Id, model.ArticleStatusPublish); page.TotalCount == 0 {
		return
	}
	page.List = dao.PageArticleByCategoryIdAndStatus(category.Id, page.PageNo, page.PageSize, model.ArticleStatusPublish)
}

func PagePublishArticle(page *model.Page) {
	if page.TotalCount = dao.CountArticleByStatus(model.ArticleStatusPublish); page.TotalCount == 0 {
		return
	}
	page.List = dao.PageArticleByStatus(page.PageNo, page.PageSize, model.ArticleStatusPublish)
}

func ListRecentlyPublishArticle(limit int) (articleList []model.Article) {
	articleList = dao.ListRecentlyArticle(model.ArticleStatusPublish, limit)
	return
}
