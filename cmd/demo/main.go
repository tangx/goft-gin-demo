package main

import (
	"github.com/shenyisyn/goft-gin/goft"
	"github.com/tangx/goft-gin-demo/pkg/controllers"
)

func main() {
	//初始化脚手架
	s := goft.Ignite()

	//挂载控制器 group v1
	s.Mount("/demo/v1", controllers.NewIndex())

	// 挂载控制器 group v2 ok
	// s.Mount("/demo/v2", controllers.NewIndexController())

	// 启动
	s.Launch() //启动 默认8080
}
