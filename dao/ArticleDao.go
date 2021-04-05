package dao

import (
	"go-blog/config"
	"go-blog/model"
)

func GetArticleById(Id *int) (article model.Article) {

	config.Db.Raw("select ID,post_content FROM wp_posts WHERE ID = ? ", Id).Scan(&article)
	return
}
