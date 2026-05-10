package tenant

import (
	"context"
	"goframe-vben/internal/service"

	"goframe-vben/api/tenant/admin"
)

func (c *ControllerAdmin) Delete(ctx context.Context, req *admin.DeleteReq) (res *admin.DeleteRes, err error) {
	return nil, service.Tenant().Delete(ctx, req.Id)
}
