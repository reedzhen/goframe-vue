package tenant_package

import (
	"context"
	"goframe-vben/internal/service"

	"goframe-vben/api/tenant_package/admin"
)

func (c *ControllerAdmin) GetMenuList(ctx context.Context, req *admin.GetMenuListReq) (res *admin.GetMenuListRes, err error) {
	list, err := service.TenantPackage().GetMenuList(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return (*admin.GetMenuListRes)(&list), nil
}
