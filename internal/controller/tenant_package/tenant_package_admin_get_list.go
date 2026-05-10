package tenant_package

import (
	"context"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"
	"strings"

	"goframe-vben/api/tenant_package/admin"
)

func (c *ControllerAdmin) GetList(ctx context.Context, req *admin.GetListReq) (res *admin.GetListRes, err error) {
	list, err := service.TenantPackage().GetList(ctx, dto.TenantPackageGetListInput{
		Name:   strings.TrimSpace(req.Name), // 套餐名
		Status: req.Status,                  // 租户状态 1正常/2停用
	})
	if err != nil {
		return nil, err
	}

	return &admin.GetListRes{List: list}, nil
}
