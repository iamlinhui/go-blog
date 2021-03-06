package dao

import (
	"blog/server/config"
	"blog/server/model"
)

func GetCategoryByPrefix(prefix string) (category model.Category) {
	if err := config.Db.Where("prefix = ? ", prefix).First(&category).Error; err != nil {
		panic(err)
	}
	return
}

func ListListCategory() (categoryList []model.Category) {
	if err := config.Db.Find(&categoryList).Order("sort").Error; err != nil {
		panic(err)
	}
	return
}
