package main

import (
	"github.com/shenyisyn/goft-gin/goft"
	"github.com/tangx/goft-gin-demo/pkg/controllers"
)

func main() {
	goft.Ignite(). //初始化脚手架
			Mount("/demo/v1", controllers.NewIndexController()). //挂载控制器
			Launch()                                             //启动 默认8080
}
