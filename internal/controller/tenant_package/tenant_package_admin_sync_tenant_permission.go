package tenant_package

import (
	"context"
	"goframe-vben/internal/service"

	"goframe-vben/api/tenant_package/admin"
)

func (c *ControllerAdmin) SyncTenantPermission(ctx context.Context, req *admin.SyncTenantPermissionReq) (res *admin.SyncTenantPermissionRes, err error) {
	return nil, service.TenantPackage().SyncTenantPermission(ctx, req.Id)
}
