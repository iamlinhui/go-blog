package model

type User struct {
	Id         uint
	Login      string    // 登录的名称 不可修改 唯一
	Email      string    // 邮箱 唯一 可修改
	Salt       string    // 盐
	Password   string    // 密码
	Name       string    // 中文名
	Status     uint      // 状态 0冻结 1可用
	Role       uint      // 0管理,1会员
	CreateTime LocalTime // 注册日期
	LoginTime  LocalTime // 最近登录日期
}
