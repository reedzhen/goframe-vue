package admin

import (
	"context"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/role"
)

func (c *ControllerRole) ChangeDataScope(ctx context.Context, req *role.ChangeDataScopeReq) (res *role.ChangeDataScopeRes, err error) {
	return nil, service.Role().ChangeDataScope(ctx, req.Id, int(req.DataScope))
}
