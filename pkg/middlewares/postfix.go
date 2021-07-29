package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Postfix struct {
	msg string
}

func NewPostfix(msg string) *Postfix {
	return &Postfix{
		msg: msg,
	}
}

func (p *Postfix) OnRequest(*gin.Context) error {
	return nil
}

func (p *Postfix) OnResponse(result interface{}) (interface{}, error) {
	// 由于 result 是 interface{}, 为了要对其进行处理， 需要进行断言获取真实类型。
	if r, ok := result.(string); ok {
		r = fmt.Sprintf("%s - by %s", r, p.msg)
		return r, nil
	}
	return result, nil
}
