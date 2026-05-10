package router

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/net/ghttp"
	"goframe-vben/addons/example/controller/user"
	"goframe-vben/internal/consts"
)

// Admin 后台路由
func Admin(ctx context.Context, group *ghttp.RouterGroup, addonName string) {
	prefix := fmt.Sprintf("/%s/%s", consts.GroupAdminApi, addonName)
	group.Group(prefix, func(group *ghttp.RouterGroup) {
		group.Bind(
			user.NewAdmin(),
		)
	})
}
