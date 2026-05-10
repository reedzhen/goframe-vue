package tenant_package

import (
	"context"
	"goframe-vben/internal/service"

	"goframe-vben/api/tenant_package/admin"
)

func (c *ControllerAdmin) Delete(ctx context.Context, req *admin.DeleteReq) (res *admin.DeleteRes, err error) {
	return nil, service.TenantPackage().Delete(ctx, req.Id)
}
