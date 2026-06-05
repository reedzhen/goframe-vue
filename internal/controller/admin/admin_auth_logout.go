package admin

import (
	"context"
	"github.com/goflyfox/gtoken/v2/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"goframe-vben/internal/consts"
	"goframe-vben/internal/library/contexts"
	"goframe-vben/internal/library/gtoken/backend"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/auth"
)

func (c *ControllerAuth) Logout(ctx context.Context, req *auth.LogoutReq) (res *auth.LogoutRes, err error) {
	// 登出销毁Token
	if err = backend.InitGToken().Destroy(ctx, g.RequestFromCtx(ctx).GetCtxVar(gtoken.KeyUserKey).String()); err != nil {
		return
	}
	user := contexts.GetUser(ctx)
	if user != nil {
		service.LoginRecord().AsyncCreate(ctx, dto.LoginRecordCreateInput{
			Username:  user.Username,
			LoginType: consts.LoginRecordTypeOut,
			Remark:    "退出登录",
		})
	}
	return
}
