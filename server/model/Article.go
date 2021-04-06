package model

import (
	"time"
)

type Article struct {
	Id            uint        `gorm:"column:id;type:int;not null" json:"id"`       // 主键
	Content       string      `gorm:"column:content;type:longtext" json:"content"` // 正文
	Title         string      // 题目
	Excerpt       string      // 摘要
	Url           string      //文章访问链接
	UserId        uint        // 作者ID
	CreateTime    time.Time   // 发布日期
	ModifyTime    time.Time   // 更新日期
	Status        uint        // 文章状态 0:未发布,1:已发布
	CommentStatus uint        // 评论开关 0:不开启,1:开启
	CommentCount  string      // 评论数量
	Category      []*Category `gorm:"-"` // 文章所属分类
}
