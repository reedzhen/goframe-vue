package admin

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"goframe-vben/api/admin/common"
)

func (c *ControllerCommon) GetRoutesWithMeta(ctx context.Context, req *common.GetRoutesWithMetaReq) (res *common.GetRoutesWithMetaRes, err error) {
	var (
		r         = g.RequestFromCtx(ctx)
		routes    = r.Server.GetRoutes()                   // 获取所有路由信息
		currRoute = r.GetServeHandler().Handler.Router.Uri // 当前请求路由 例如：/admin-api/auth/user
	)

	list := make([]*common.RoutesWithMetaItem, 0)
	for _, route := range routes {
		// 跳过一些特殊路由，如中间件、钩子函数等
		if !route.IsServiceHandler {
			continue
		}
		// 获取路由的元数据信息
		var (
			tags    = route.Handler.GetMetaTag("tags")    // 标签/分组
			summary = route.Handler.GetMetaTag("summary") // 简要描述
			path    = route.Handler.GetMetaTag("path")
			method  = route.Handler.GetMetaTag("method")
		)
		if tags == "" {
			continue
		}
		// 补上前缀
		if gstr.HasPrefix(currRoute, "/admin-api") {
			path = "/admin-api" + path
		}
		list = append(list, &common.RoutesWithMetaItem{
			Tags:    tags,
			Summary: summary,
			Path:    path,
			Method:  method,
		})
	}
	return (*common.GetRoutesWithMetaRes)(&list), nil
}
