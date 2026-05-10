package tenant

import (
	"context"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"
	"strings"

	"goframe-vben/api/tenant/admin"
)

func (c *ControllerAdmin) GetPage(ctx context.Context, req *admin.GetPageReq) (res *admin.GetPageRes, err error) {
	data, err := service.Tenant().GetPage(ctx, dto.TenantPageInput{
		PageParam:  req.PageParam,
		TenantName: strings.TrimSpace(req.TenantName), // 租户名称
		LinkMan:    strings.TrimSpace(req.LinkMan),    // 联系人姓名
		LinkPhone:  strings.TrimSpace(req.LinkPhone),  // 联系人手机
		Status:     req.Status,                        // 租户状态 1正常/2停用
	})
	return &admin.GetPageRes{Result: data}, err
}
