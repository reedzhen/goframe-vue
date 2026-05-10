package router

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/net/ghttp"
	"goframe-vben/internal/consts"
)

// OpenApi 开放给第三方的接口
func OpenApi(ctx context.Context, group *ghttp.RouterGroup) {
	//gfTokenOpen := gtoken.OpenStartGToken()
	group.Group(fmt.Sprintf("/%s", consts.GroupOpenApi), func(group *ghttp.RouterGroup) {
		//group.Bind(
		//	auth.NewOpen(),
		//).Middleware(service.Middleware().CheckSum)

		//_ = gfTokenOpen.Middleware(ctx, group)
		//g.Log().Debug(ctx, fmt.Sprintf("🎉 [GToken] %s start... ", consts.GroupOpenApi))
		//
		//group.Bind(
		//	user.NewOpen(), // 用户管理
		//)
	})
}
