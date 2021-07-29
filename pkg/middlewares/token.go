package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
)

// TokenCheck 定义对象
// 并创建 OnRequest 与 OnResponse 方法， 以满足 goft.Fairing 接口
//	https://github.com/shenyisyn/goft-gin/blob/v0.5.2/goft/Fairing.go#L5
type TokenCheck struct{}

// NewTokenCheck 初始化中间件
func NewTokenCheck() *TokenCheck {
	return &TokenCheck{}
}

// OnRequest 为了满足 Fairing Middleware 接口
// 	https://github.com/shenyisyn/goft-gin/blob/v0.5.2/goft/FairingHandler.go#L33
func (tc *TokenCheck) OnRequest(c *gin.Context) error {
	defer func() {
		fmt.Println("this will be printed, even goft.Throw trigger a panic")
	}()

	if c.Query("token") == "" {
		// goft.Throw 1. 引发一个 panic， 最终将在被捕获并向用户返回异常
		// 	2. 定义错误返回状态码, ex: 503
		// 	3. 满足 panic 行为
		goft.Throw("token required", 503, c)
	}
	return nil
}

// OnResponse 为了满足 Fairing Middleware 接口
// 	https://github.com/shenyisyn/goft-gin/blob/v0.5.2/goft/FairingHandler.go#L41
func (tc *TokenCheck) OnResponse(result interface{}) (interface{}, error) {
	if r, ok := result.(string); ok {
		r = fmt.Sprintf("%s - by token-check", r)
		return r, nil
	}
	return result, nil
}
