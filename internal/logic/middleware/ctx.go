package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"goframe-vben/internal/boot"
	"goframe-vben/internal/library/contexts"
	"goframe-vben/internal/model"
)

// Ctx 请求结束时不会自动调用Done方法结束掉ctx
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	// creates and returns a never done context object
	r.SetCtx(r.GetNeverDoneCtx())

	// 初始化上下文
	contexts.Init(r, &model.Context{
		Module: boot.GetCurrentModule(r),
	})

	// 上下文中记录路由信息，方便记录操作日志时获取
	setRouteInfo(r)

	r.Middleware.Next()
}

// setRouteInfo 上下文中写入当前规范路由 tags 和 summary
func setRouteInfo(r *ghttp.Request) {
	handler := r.GetServeHandler()
	if handler != nil {
		contexts.SetData(r.Context(), g.Map{
			"tags":    handler.GetMetaTag("tags"),
			"summary": handler.GetMetaTag("summary"),
		})
	}
}
