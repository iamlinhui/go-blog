package model

type Exception struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

const (
	Success  = 0
	Error    = 100000
	ErrorId  = 200000
	NotFound = 200001
)

var codeMsg = map[int]string{
	Success:  "OK",
	Error:    "FAIL",
	ErrorId:  "输入ID不合法",
	NotFound: "数据不存在",
}

func GetException(code int) *Exception {
	return &Exception{
		Code:    code,
		Message: codeMsg[code],
	}
}
