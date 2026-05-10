package tenant_package

import (
	"context"
	"goframe-vben/internal/service"

	"goframe-vben/api/tenant_package/admin"
)

func (c *ControllerAdmin) GetInfo(ctx context.Context, req *admin.GetInfoReq) (res *admin.GetInfoRes, err error) {
	data, err := service.TenantPackage().GetInfo(ctx, req.Id)
	if err != nil {
		return
	}

	return &admin.GetInfoRes{Data: data}, nil
}
