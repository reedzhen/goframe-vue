package tenant_package

import (
	"context"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"
	"strings"

	"goframe-vben/api/tenant_package/admin"
)

func (c *ControllerAdmin) GetPage(ctx context.Context, req *admin.GetPageReq) (res *admin.GetPageRes, err error) {
	data, err := service.TenantPackage().GetPage(ctx, dto.TenantPackagePageInput{
		PageParam: req.PageParam,
		Name:      strings.TrimSpace(req.Name), // 套餐名
		Status:    req.Status,                  // 租户状态 1正常/2停用
	})
	return &admin.GetPageRes{Result: data}, err
}
