package tenant_package

import (
	"context"
	"goframe-vben/internal/library/contexts"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"
	"strings"

	"goframe-vben/api/tenant_package/admin"
)

func (c *ControllerAdmin) Update(ctx context.Context, req *admin.UpdateReq) (res *admin.UpdateRes, err error) {
	return nil, service.TenantPackage().Update(ctx, dto.TenantPackageUpdateInput{
		Id:        req.Id,
		UpdatedBy: contexts.GetUserId(ctx),
		TenantPackageCreateUpdateBase: dto.TenantPackageCreateUpdateBase{
			Name:    strings.TrimSpace(req.Name),    // 套餐名
			Status:  req.Status,                     // 租户状态 1正常/2停用
			MenuIds: strings.TrimSpace(req.MenuIds), // 菜单ID 逗号分隔
			Remark:  strings.TrimSpace(req.Remark),  // 备注
		},
	})
}
