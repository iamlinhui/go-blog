package model

type Category struct {
	Id      uint
	Name    string // 中文名(展示)
	Prefix  string // 访问前缀(url)
	Checked bool   //是否选中
}
