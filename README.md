# goft-gin demo


demo 的目录结构将尽量与 https://github.com/golang-standards/project-layout 保持一致

## Documents and Usage

0. [Config file](./cmd/demo/application.yaml)
1. Goft [Initial and Launch](./cmd/demo/main.go)
2. Controller [Route and Handler](./pkg/controllers/index.go)
3. Fairing [Middleware](./pkg/middlewares/token.go)

如果使用了 `Fairing Middleware` , 那么数据的流转将会按照一下过程执行。

```go
user
    -> gin -> request
        -> Fairing.OnRequest(c *gin.Context) error
    -> Controller.HandlerFunc -> response
        -> Fairing.OnResponse(result interface{})(interface{},error) -> response
-> user
```
