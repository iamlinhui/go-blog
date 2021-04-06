package dao

import (
	"blog/src/config"
	"blog/src/model"
)

func GetArticleById(Id *int) (article model.Article) {
	if err := config.Db.Raw("SELECT ID AS id,post_content FROM wp_posts WHERE ID = ? ", Id).First(&article).Error; err != nil {
		panic(err)
	}
	return
}
