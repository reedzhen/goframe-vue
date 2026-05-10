package contexts

import (
	"context"
	"goframe-vben/internal/consts"
	"goframe-vben/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// Init 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
func Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(consts.ContextHTTPKey, customCtx)
}

// Get 获得上下文变量，如果没有设置，那么返回nil
func Get(ctx context.Context) *model.Context {
	value := ctx.Value(consts.ContextHTTPKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}

// GetUser 获取用户信息
func GetUser(ctx context.Context) *model.Identity {
	c := Get(ctx)
	if c == nil {
		return nil
	}
	return c.User
}

// GetUserId 获取用户Id
func GetUserId(ctx context.Context) int64 {
	user := GetUser(ctx)
	if user == nil {
		return 0
	}
	return user.UserId
}

// IsAdmin 是否是超管
func IsAdmin(ctx context.Context) bool {
	user := GetUser(ctx)
	if user == nil {
		return false
	}
	return user.IsAdmin
}

// GetRoleId 获取用户角色Id
func GetRoleId(ctx context.Context) int64 {
	user := GetUser(ctx)
	if user == nil {
		return 0
	}

	return user.RoleId
}

// GetOrganizationId 获取获取组织Id
func GetOrganizationId(ctx context.Context) int64 {
	user := GetUser(ctx)
	if user == nil {
		return 0
	}
	return user.OrganizationId
}

// GetModule 获取应用模块 admin/api/open
func GetModule(ctx context.Context) string {
	c := Get(ctx)
	if c == nil {
		return ""
	}
	return c.Module
}

// GetData 获取额外数据
func GetData(ctx context.Context) g.Map {
	c := Get(ctx)
	if c == nil {
		return nil
	}
	return c.Data
}

// SetUser 将用户信息设置到上下文请求中，注意是完整覆盖
func SetUser(ctx context.Context, ctxUser *model.Identity) {
	c := Get(ctx)
	if c == nil {
		g.Log().Warning(ctx, "contexts.SetUser, c == nil ")
		return
	}
	c.User = ctxUser
}

// SetData 设置额外数据
func SetData(ctx context.Context, data g.Map) {
	Get(ctx).Data = data
}

// SetAddonName 设置插件信息
func SetAddonName(ctx context.Context, name string) {
	c := Get(ctx)
	if c == nil {
		g.Log().Warning(ctx, "contexts.SetAddonName, c == nil ")
		return
	}
	Get(ctx).AddonName = name
}
