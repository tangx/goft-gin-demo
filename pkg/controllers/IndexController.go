package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
)

type IndexController struct{}

func NewIndexController() *IndexController {
	return &IndexController{}
}
func (this *IndexController) Index(ctx *gin.Context) string {
	return "this is 首页"
}
func (this *IndexController) Name() string {
	return "IndexController"
}
func (this *IndexController) Build(goft *goft.Goft) {
	goft.Handle("GET", "/", this.Index)
}
