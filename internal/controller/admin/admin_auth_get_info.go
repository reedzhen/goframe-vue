package admin

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-vben/internal/codes"
	"goframe-vben/internal/library/contexts"
	"goframe-vben/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"

	"goframe-vben/api/admin/auth"
)

func (c *ControllerAuth) GetInfo(ctx context.Context, req *auth.GetInfoReq) (res *auth.GetInfoRes, err error) {
	info, err := service.User().ValidateExists(ctx, contexts.GetUserId(ctx))
	if err != nil {
		return
	}
	if info == nil {
		// 返回 401，前端识别code码重新登录
		return nil, gerror.NewCode(codes.CodeUnauthorized, "用户不存在")
	}
	if err = gconv.Struct(info, &res); err != nil {
		return
	}
	res.RealName = info.Nickname

	// 菜单和权限
	if res.Authorities, err = service.Menu().GetAuthorizedList(ctx); err != nil {
		return
	}

	return
}
