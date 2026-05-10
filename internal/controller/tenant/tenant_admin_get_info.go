package tenant

import (
	"context"
	"goframe-vben/internal/service"

	"goframe-vben/api/tenant/admin"
)

func (c *ControllerAdmin) GetInfo(ctx context.Context, req *admin.GetInfoReq) (res *admin.GetInfoRes, err error) {
	data, err := service.Tenant().GetInfo(ctx, req.Id)
	if err != nil {
		return
	}

	return &admin.GetInfoRes{Data: data}, nil
}
