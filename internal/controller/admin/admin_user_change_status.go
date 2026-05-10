package admin

import (
	"context"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/user"
)

func (c *ControllerUser) ChangeStatus(ctx context.Context, req *user.ChangeStatusReq) (res *user.ChangeStatusRes, err error) {
	return nil, service.User().ChangeStatus(ctx, req.UserId, int(req.Status))
}
