package errors

import "fmt"

// 定义自定义错误结构体
type CustomError struct {
	Code    int
	Message string
}

// 实现 error 接口的 Error 方法
func (e *CustomError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

// 构造函数
func New(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

// 预定义错误示例
var (
	ErrNotFound   = New(404, "资源未找到")
	ErrInternal   = New(500, "内部服务器错误")
	ErrBadRequest = New(400, "错误的请求")
)
