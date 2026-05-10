package router

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/net/ghttp"
	"goframe-vben/internal/consts"
	"goframe-vben/internal/controller/admin"
	"goframe-vben/internal/controller/tenant"
	"goframe-vben/internal/controller/tenant_package"
	"goframe-vben/internal/service"
)

// AdminApi 后端接口
func AdminApi(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group(fmt.Sprintf("/%s", consts.GroupAdminApi), func(group *ghttp.RouterGroup) {
		// 全局中间件
		group.Middleware(
			service.Middleware().GlobalRateLimit, // 全局IP限流：200次/分钟
			service.Middleware().CheckJwtToken,   // gtoken检验
			service.Middleware().CheckApiAuth,    // 接口权限校验
		)
		group.Bind(
			admin.NewCommon(),          // 通用接口
			admin.NewAuth(),            // 授权
			admin.NewRole(),            // 角色
			admin.NewUser(),            // 用户
			admin.NewMenu(),            // 菜单
			admin.NewLoginrecord(),     // 登录日志
			admin.NewConfig(),          // 配置
			admin.NewDict(),            // 字典类型
			admin.NewOperationrecord(), // 操作日志
			admin.NewCron(),            // 定时任务
			admin.NewOrganization(),    // 组织管理
			tenant.NewAdmin(),          // 贴牌管理
			tenant_package.NewAdmin(),  // 贴牌套餐
		)
	})
}
