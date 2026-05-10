package admin

import (
	"context"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/role"
)

func (c *ControllerRole) Delete(ctx context.Context, req *role.DeleteReq) (res *role.DeleteRes, err error) {
	return nil, service.Role().Delete(ctx, req.Id)
}
