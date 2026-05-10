package admin

import (
	"context"
	"goframe-vben/internal/codes"
	"goframe-vben/internal/consts"
	"goframe-vben/internal/library/captcha"
	"goframe-vben/internal/library/gtoken/backend"
	"goframe-vben/internal/library/ratelimit"

	"github.com/gogf/gf/v2/errors/gerror"
	"goframe-vben/api/admin/auth"
	"goframe-vben/internal/service"
)

func (c *ControllerAuth) Login(ctx context.Context, req *auth.LoginReq) (res *auth.LoginRes, err error) {
	// 1. 限流检查：防止暴力破解（每分钟最多5次）
	if err = ratelimit.CheckLoginLimit(ctx); err != nil {
		return nil, err
	}

	// 2. 校验验证码
	if !captcha.VerifyAndClear(ctx, req.CaptchaKey, req.CaptchaCode) {
		return nil, gerror.NewCode(codes.CodeAuthCaptchaError)
	}

	// 3. 获取用户信息
	user, err := service.User().GetInfoByUsername(ctx, req.Username)
	if err != nil {
		return
	}
	if user == nil {
		return nil, gerror.NewCode(codes.CodeAuthCredentialError)
	}
	if user.Status != consts.UserStatusOk {
		return nil, gerror.NewCode(codes.CodeAuthUserFrozen)
	}

	// 4. 获取角色
	roleInfo, err := service.Role().ValidateExists(ctx, user.RoleId)
	if err != nil {
		return
	}
	// 判断角色状态
	if roleInfo.Status != int(consts.RoleStatusOk) {
		return nil, gerror.NewCode(codes.CodeAuthRoleDisabled)
	}

	// 5. 生成Token
	token, err := backend.InitGToken().Generate(ctx, user.Username, user)
	if err != nil {
		return
	}

	return &auth.LoginRes{
		AccessToken: token,
	}, nil
}
