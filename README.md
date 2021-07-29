# goft-gin demo


demo 的目录结构将尽量与 https://github.com/golang-standards/project-layout 保持一致

## Documents and Usage

0. [Config file](./cmd/demo/application.yaml)
1. Goft [Initial and Launch](./cmd/demo/main.go)
    + [config] 使用 `goft.Config` 注入 **配置**
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


4. Inject [Config](./pkg/config/db_config.go)
    + 配置在 `goft.Config` 中完成 IoC 容器的配置声明与注入
    + 在 Controller 配置中使用 tag `inject:""` 获取注入配置 [controllers/User.go demo #L12](./pkg/controllers/user.go#L12): 1. 字段名字随意但必须公开，2. 类型必须 main.Config

5. Inject [Db Config and Adapter #42 ](./pkg/config/db_config.go#L42)
    + [db] 如果要支持 `goft.SimpleQuery / goft.Query` 那么 db 需要支持方法签名 `func (db *driver.DB) DB() *sql.DB)`
    + [db] 否则就需要自己构造 **Adapter** 实现上述签名。
