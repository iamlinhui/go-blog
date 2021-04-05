package dao

import (
	"blog/src/config"
	"blog/src/model"
)

func GetArticleById(Id *int) (article model.Article) {

	config.Db.Raw("select ID,post_content FROM wp_posts WHERE ID = ? ", Id).Scan(&article)
	return
}
