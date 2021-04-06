package dao

import (
	"blog/server/config"
	"blog/server/model"
)

func GetArticleById(id int) (article model.Article) {
	if err := config.Db.Where("id = ? ", id).First(&article).Error; err != nil {
		panic(err)
	}
	return
}

func GetArticleByUrl(url string) (article model.Article) {
	if err := config.Db.Where("url = ? ", url).First(&article).Error; err != nil {
		panic(err)
	}
	return
}

func PageArticleByCategoryId(categoryId uint, pageNo int, pageSize int, status int) (articleList []model.Article) {
	sql := `
			SELECT id,
				   content,
				   title,
				   excerpt,
				   url,
				   user_id,
				   create_time,
				   modify_time,
				   status,
				   comment_status
			FROM t_article
					 LEFT JOIN t_relationships ON id = article_id
			WHERE category_id = ?
			  AND status = ?
			LIMIT ?,?
			`
	if err := config.Db.Raw(sql, categoryId, status, (pageNo-1)*pageSize, pageSize).Scan(&articleList).Error; err != nil {
		panic(err)
	}
	return
}

func CountArticleByCategoryId(categoryId uint, status int) (count int) {
	sql := `
			SELECT COUNT(*)
			FROM t_article
					 LEFT JOIN t_relationships ON id = article_id
			WHERE category_id = ?
			  AND status = ?
			`
	if err := config.Db.Raw(sql, categoryId, status).Scan(&count).Error; err != nil {
		panic(err)
	}
	return
}
