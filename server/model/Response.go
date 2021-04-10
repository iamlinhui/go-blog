package model

type Response struct {
	Exception
	Data interface{} `json:"data"`
}

func Ok(data interface{}) *Response {
	return &Response{
		Exception: Exception{
			Code:    Success,
			Message: codeMsg[Success],
		},
		Data: data,
	}
}

func Fail(code int, data interface{}) *Response {
	return &Response{
		Exception: Exception{
			Code:    code,
			Message: codeMsg[code],
		},
		Data: data,
	}
}
