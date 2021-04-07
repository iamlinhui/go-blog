package service

import (
	"blog/server/dao"
	"blog/server/model"
)

func ListCategory() (categoryList []model.Category) {
	categoryList = dao.ListListCategory()
	return
}
