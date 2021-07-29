package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
)

type IndexController struct{}

func NewIndexController() *IndexController {
	return &IndexController{}
}

// Name 为了向 goft.Goft.exprData 中注册 Controller。 需要全局唯一。
// 		https://github.com/shenyisyn/goft-gin/blob/3e3f783147166ca2c3a7c14ac9aecce46bdeeaed/goft/Goft.go#L107
func (ic *IndexController) Name() string {
	return "IndexController"
}

// Build 1. 注册路由地址 2. 为路由绑定 Handler
//	1. https://github.com/shenyisyn/goft-gin/blob/3e3f783147166ca2c3a7c14ac9aecce46bdeeaed/goft/Goft.go#L128
// 	2. https://github.com/shenyisyn/goft-gin/blob/3e3f783147166ca2c3a7c14ac9aecce46bdeeaed/goft/Goft.go#L61
func (ic *IndexController) Build(goft *goft.Goft) {
	goft.Handle("GET", "/index", ic.handlerIndex)
	goft.Handle("GET", "/json", ic.handlerIndexJson)
}

// handlerIndex  类似 gin.HandlerFunc, 请求处理
// 	 handler 返回 interface{}, 可以返回任意值。
func (ic *IndexController) handlerIndex(ctx *gin.Context) string {
	return "this is 首页"
}

// handlerIndexJson 使用了 goft 默认返回中的其一， json
// 	默认返回类型: 1. Json, 2. Query 3. SimpleQuery 4. Void
// 		5. View 预留，但暂时不支持
// 	参考 goft/Responder.go 文件
func (ic *IndexController) handlerIndexJson(ctx *gin.Context) goft.Json {
	return map[string]string{"json": "hanlder"}
}
