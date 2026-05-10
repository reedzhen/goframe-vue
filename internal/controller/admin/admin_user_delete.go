package admin

import (
	"context"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/user"
)

func (c *ControllerUser) Delete(ctx context.Context, req *user.DeleteReq) (res *user.DeleteRes, err error) {
	return nil, service.User().Delete(ctx, req.Id)
}
