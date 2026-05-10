package plugin

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/os/gctx"
	"goframe-vben/internal/library/addons"
	"goframe-vben/internal/router"
	"goframe-vben/internal/service"
	"goframe-vben/utility/response"
)

type Http struct {
}

func (p *Http) Name() string {
	return "http"
}

func (p *Http) Author() string {
	return "Larry Liu"
}

func (p *Http) Version() string {
	return "1.0"
}

func (p *Http) Description() string {
	return "http service"
}

func (p *Http) Install(s *ghttp.Server) error {
	ctx := gctx.GetInitCtx()

	s.BindStatusHandler(404, func(r *ghttp.Request) {
		r.Response.Writeln("woops, status 404 found")
	})

	// test
	s.BindHandler("/hello", func(r *ghttp.Request) {
		response.JsonExit(r, gcode.CodeOK.Code(), "", "world")
	})

	//// queue test
	//s.BindHandler("/ping", func(r *ghttp.Request) {
	//	g.Dump(asynqueue.Client.EnqueueContext(ctx, task.NewPingTask("ping"), asynq.TaskID("ping_123")))
	//})

	s.Group("/", func(group *ghttp.RouterGroup) {

		// 注册全局中间件
		group.Middleware(
			service.Middleware().CORS, // 跨域中间件，自动处理跨域问题
			service.Middleware().Ctx,  // 初始化请求上下文
			//service.Middleware().SetTenant,       // 设置租户信息
			service.Middleware().OperationRecord, // 记录操作日志
			service.Middleware().HandlerResponse, // HTTP响应处理，在业务处理完成后，对响应结果进行格式化和错误过滤，将处理后的数据发送给请求方
		)

		// 注册后台路由
		router.AdminApi(ctx, group)

		// 注册前端路由
		router.AppApi(ctx, group)

		// 注册第三方路由
		router.OpenApi(ctx, group)

		// 注册插件路由
		addons.RegisterModulesRouter(ctx, group)
	})

	// OpenApi自定义信息
	oai := s.GetOpenApi()
	oai.Info.Title = `API 文档`
	oai.Info.Description = `
## 全局错误码说明

| 错误码 | 消息内容 | 说明 |
| :--- | :--- | :--- |
| **0** | **操作成功** | 请求处理正常 |
| **400** | **请求参数不正确** | 客户端请求参数校验失败 |
| **401** | **账号未登录** | Token 失效或未带 Authorization 报文头 |
| **403** | **没有该操作权限** | 当前账号权限等级不够 |
| **404** | **请求未找到** | 请求的路由或资源不存在 |
| **405** | **请求方法不支持** | HTTP Method 与接口定义不匹配 |
| **423** | **请求失败，请稍后重试** | 资源锁定或并发冲突 |
| **429** | **请求过于频繁** | 触发限流，请稍后重试 |
| **500** | **系统异常** | 服务器内部代码执行或数据库异常 |
| **501** | **正在开发中** | 接口已定义但还未完成具体业务逻辑 |
| **900** | **重复请求** | 接口幂等性校验未通过 |
| **901** | **演示模式** | 系统处于演示模式，已禁止写操作 |
| **999** | **未知错误** | 未捕获的系统底层错误 |

> **提示**：除 0 以外的错误码均表示请求异常，对应的错误信息会通过 message 字段返回。
`
	oai.Config.CommonResponse = response.JsonRes{}
	oai.Config.CommonResponseDataField = `Data`
	// 手动定义 Tags 的展示顺序
	oai.Tags = &goai.Tags{
		{Name: "通用接口", Description: "后台通用接口"},
		{Name: "字典管理", Description: ""},
		{Name: "用户登录", Description: "系统权限与认证相关接口"},
		{Name: "用户管理", Description: ""},
		{Name: "角色管理", Description: ""},
		{Name: "菜单管理", Description: ""},
		{Name: "系统配置", Description: ""},
		{Name: "定时任务", Description: ""},
		{Name: "组织机构", Description: ""},
		{Name: "登录日志", Description: ""},
		{Name: "操作日志", Description: ""},
	}

	return nil
}

func (p *Http) Remove() error {
	return nil
}
