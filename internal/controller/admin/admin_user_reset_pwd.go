package admin

import (
	"context"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/user"
)

func (c *ControllerUser) ResetPwd(ctx context.Context, req *user.ResetPwdReq) (res *user.ResetPwdRes, err error) {
	return nil, service.User().ResetPwd(ctx, req.UserId, req.Password)
}
