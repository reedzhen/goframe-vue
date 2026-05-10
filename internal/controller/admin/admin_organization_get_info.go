package admin

import (
	"context"
	"goframe-vben/api/admin/organization"
	"goframe-vben/internal/service"
)

// GetInfo 详情
func (c *ControllerOrganization) GetInfo(ctx context.Context, req *organization.GetInfoReq) (res *organization.GetInfoRes, err error) {
	data, err := service.Organization().GetInfo(ctx, req.Id)
	if err != nil {
		return
	}

	return &organization.GetInfoRes{SysOrganization: data}, nil
}
