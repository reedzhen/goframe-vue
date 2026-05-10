package admin

import (
	"context"
	"github.com/goflyfox/gtoken/v2/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"goframe-vben/internal/library/gtoken/backend"

	"goframe-vben/api/admin/auth"
)

func (c *ControllerAuth) Logout(ctx context.Context, req *auth.LogoutReq) (res *auth.LogoutRes, err error) {
	// 登出销毁Token
	if err = backend.InitGToken().Destroy(ctx, g.RequestFromCtx(ctx).GetCtxVar(gtoken.KeyUserKey).String()); err != nil {
		return
	}
	return
}
