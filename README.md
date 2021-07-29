# goft-gin demo


demo 的目录结构将尽量与 https://github.com/golang-standards/project-layout 保持一致

## Documents and Usage

0. [Config file](./cmd/demo/application.yaml)
1. Goft [Initial and Launch](./cmd/demo/main.go)
    + [`middleware`] 使用 `goft.Attach` 绑定 **全局路由 Fairing middleware**
    + [`api`] 使用 `goft.Mount` 挂载 **路由组**
2. Controller(`RESTful API`) [Route and Handler](./pkg/controllers/index.go)
    + [`api`] 使用 `goft.Handle(...)` 注册 **相对路由**
3. Fairing(`Middleware`) [Handle](./pkg/middlewares/token.go)
    + [`middleware`] 使用 `goft.HandleWithFairing(...)` 绑定 **局部路由 Fairing middleware**

> 注意: 使用 `goft.Handle(,"/path1",).HandleWithFairing(,"/path2",)` 将会创建两条 **同级** 相对路由。
> 即将产生 **两条路由** `/v1/path1, /v1/path2` 而非 **一条路由** `/v1/path1/path2`。

如果使用了 `Fairing Middleware` , 那么数据的流转将会按照一下过程执行。

```go
user
    -> gin -> request
        -> Fairing.OnRequest(c *gin.Context) error
    -> Controller.HandlerFunc -> response
        -> Fairing.OnResponse(result interface{})(interface{},error) -> response
-> user
```
