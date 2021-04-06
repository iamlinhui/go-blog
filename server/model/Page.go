package model

type Page struct {
	TotalCount int         `json:"totalCount" `
	Data       interface{} `json:"data"`
	PageNo     int         `json:"pageNo" binding:"required"`
	PageSize   int         `json:"pageSize" binding:"required"`
}
