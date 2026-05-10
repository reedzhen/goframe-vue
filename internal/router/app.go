package router

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/net/ghttp"
	"goframe-vben/internal/consts"
)

// AppApi 前端接口
func AppApi(ctx context.Context, group *ghttp.RouterGroup) {
	//gfTokenApi := gtoken.ApiStartGToken()
	group.Group(fmt.Sprintf("/%s", consts.GroupAppApi), func(group *ghttp.RouterGroup) {
		//group.Bind(
		//	auth.NewApi(),
		//)
		//
		//_ = gfTokenApi.Middleware(ctx, group)
		//g.Log().Debug(ctx, fmt.Sprintf("🎉 [GToken] %s start... ", consts.GroupAppApi))

		//group.Bind(
		//	user.NewApi(), // 用户管理
		//)
	})
}
