package main

import (
	"github.com/shenyisyn/goft-gin/goft"
	"github.com/tangx/goft-gin-demo/pkg/config"
	"github.com/tangx/goft-gin-demo/pkg/controllers"
	"github.com/tangx/goft-gin-demo/pkg/middlewares"
)

func main() {
	// 1. 初始化脚手架
	s := goft.Ignite().Config(
		config.NewMysqlConfig(),
		config.NewServicesConfig(),
	)
	// 1.1. 初始化配置

	// 2. 注册全局中间件 middleware
	// 	2.1. 可以注册多个
	// 	2.2. 按照注册顺序执行
	// 	2.3. 相同中间件可以重复注册，并且都会执行
	s.Attach(
		middlewares.NewTokenCheck(),
		middlewares.NewPostfix("global middleware"),
		middlewares.NewTokenCheck(), // 可以重复
	)

	// 3. 挂载路由
	//	3.1. 控制器 group v1
	s.Mount("/demo/v1",
		controllers.NewIndex(),
		&controllers.User{},
		&controllers.UserInfo{},
	)
	// 	3.2. 控制器 group v2
	s.Mount("/demo/v2", controllers.NewIndex())

	// 启动
	s.Launch() //启动 默认8080
}
