package model

type Category struct {
	Id     uint   `gorm:"id" json:"id"`
	Name   string `gorm:"name" json:"name"`     // 中文名(展示)
	Prefix string `gorm:"prefix" json:"prefix"` // 访问前缀(url)
}
