package admin

import (
	"context"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/user"
)

func (c *ControllerUser) UpdatePwd(ctx context.Context, req *user.UpdatePwdReq) (res *user.UpdatePwdRes, err error) {
	return nil, service.User().UpdatePwd(ctx, req.OldPassword, req.NewPassword)
}
