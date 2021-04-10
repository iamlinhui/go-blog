package model

type Exception struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

const (
	Success     = 0
	Error       = 100000
	ErrorId     = 200000
	ErrorUrl    = 200001
	ErrorPrefix = 200002
	NotFound    = 400000
	UrlNotFound = 400404
)

var codeMsg = map[int]string{
	Success:     "OK",
	Error:       "FAIL",
	ErrorId:     "输入ID不合法",
	ErrorUrl:    "输入Url不合法",
	ErrorPrefix: "分类前缀不合法",
	NotFound:    "数据不存在",
	UrlNotFound: "请求路径不正确",
}

func GetException(code int) *Exception {
	return &Exception{
		Code:    code,
		Message: codeMsg[code],
	}
}

func BuildException(code int, msg string) *Exception {
	return &Exception{
		Code:    code,
		Message: msg,
	}
}
