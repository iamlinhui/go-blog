package model

import (
	"time"
)

type Article struct {
	Id            uint      `gorm:"id" json:"id"`                        // 主键
	Content       string    `gorm:"content" json:"content"`              // 正文
	Title         string    `gorm:"title" json:"title"`                  // 题目
	Excerpt       string    `gorm:"excerpt" json:"excerpt"`              // 摘要
	Url           string    `gorm:"url" json:"url"`                      //文章访问链接
	UserId        uint      `gorm:"user_id" json:"userId"`               // 作者ID
	CreateTime    time.Time `gorm:"create_time" json:"createTime"`       // 发布日期
	ModifyTime    time.Time `gorm:"modify_time" json:"modifyTime"`       // 更新日期
	Status        uint      `gorm:"status" json:"status"`                // 文章状态 0:未公开,1:公开
	CommentStatus uint      `gorm:"comment_status" json:"commentStatus"` // 评论开关 0:不开启,1:开启
}

const (
	ArticleStatusPublish    = 1
	ArticleStatusNotPublish = 0
)
