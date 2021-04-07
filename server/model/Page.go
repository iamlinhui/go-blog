package model

type Page struct {
	TotalCount int         `json:"totalCount" `
	Data       interface{} `json:"data"`
	PageNo     int         `json:"pageNo" binding:"required,min=1"`
	PageSize   int         `json:"pageSize" binding:"required,min=1"`
}
