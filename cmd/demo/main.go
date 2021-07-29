package main

import (
	"github.com/shenyisyn/goft-gin/goft"
	"github.com/tangx/goft-gin-demo/pkg/controllers"
	"github.com/tangx/goft-gin-demo/pkg/middlewares"
)

func main() {
	// 1. 初始化脚手架
	s := goft.Ignite()

	// 2. 注册全局中间件 middleware
	s.Attach(middlewares.NewTokenCheck())

	// 3. 挂载路由
	//	3.1. 控制器 group v1
	s.Mount("/demo/v1", controllers.NewIndex())
	// 	3.2. 控制器 group v2
	s.Mount("/demo/v2", controllers.NewIndex())

	// 启动
	s.Launch() //启动 默认8080
}
