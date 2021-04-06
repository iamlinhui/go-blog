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

func GetArticleByIdAndStatus(url string, status int) (article model.Article) {
	if err := config.Db.Where("url = ? AND status = ?", url, status).First(&article).Error; err != nil {
		panic(err)
	}
	return
}

func PageArticleByCategoryIdAndStatus(categoryId uint, pageNo int, pageSize int, status int) (articleList []model.Article) {
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

func CountArticleByCategoryIdAndStatus(categoryId uint, status int) (count int) {
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

func CountArticleByStatus(status int) (count int) {
	sql := `
			SELECT COUNT(*)
			FROM t_article
					 LEFT JOIN t_relationships ON id = article_id
			WHERE AND status = ?
			`
	if err := config.Db.Raw(sql, status).Scan(&count).Error; err != nil {
		panic(err)
	}
	return
}

func PageArticleByStatus(pageNo int, pageSize int, status int) (articleList []model.Article) {
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
			WHERE status = ?
			LIMIT ?,?
			`
	if err := config.Db.Raw(sql, status, (pageNo-1)*pageSize, pageSize).Scan(&articleList).Error; err != nil {
		panic(err)
	}
	return
}
