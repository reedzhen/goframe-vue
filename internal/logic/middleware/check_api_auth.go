package middleware

import (
	"fmt"
	"github.com/gogf/gf/v2/net/ghttp"
	"goframe-vben/internal/codes"
	"goframe-vben/internal/library/contexts"
	"strings"

	"goframe-vben/internal/service"
	"goframe-vben/utility/response"
)

// CheckApiAuth 校验接口权限中间件
func (s *sMiddleware) CheckApiAuth(r *ghttp.Request) {
	var (
		ctx     = r.GetCtx()
		handler = r.GetServeHandler()        // 获取当前路由处理函数
		route   = handler.Handler.Router.Uri // 当前请求路由
		method  = r.Method                   // 当前请求方法
	)

	// 超管直接放行
	if contexts.IsAdmin(ctx) {
		r.Middleware.Next()
		return
	}

	// 检查是否有noAuth元数据，如果有且值为true，则跳过验证
	noAuth := handler.GetMetaTag("noAuth")
	if noAuth == "true" {
		r.Middleware.Next()
		return
	}

	roleId := contexts.GetRoleId(ctx)
	if roleId == 0 {
		response.JsonExit(r, codes.CodeForbidden.Code(), codes.CodeForbidden.Message())
	}

	// 检查角色是否拥有指定权限
	hasPermission, err := service.Menu().CheckRoleApiPermission(ctx, roleId, route, method)
	if err != nil {
		response.JsonExit(r, codes.CodeForbidden.Code(), err.Error())
	}
	if !hasPermission {
		response.JsonExit(r, codes.CodeForbidden.Code(), codes.CodeForbidden.Message()+fmt.Sprintf(` ["%s:%s"]`, strings.ToLower(method), route))
		return
	}

	// 权限验证通过，继续处理请求
	r.Middleware.Next()
}
