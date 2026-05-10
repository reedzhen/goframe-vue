package tenant

import (
	"context"
	"goframe-vben/internal/library/contexts"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"
	"strings"

	"goframe-vben/api/tenant/admin"
)

func (c *ControllerAdmin) Update(ctx context.Context, req *admin.UpdateReq) (res *admin.UpdateRes, err error) {
	return nil, service.Tenant().Update(ctx, dto.TenantUpdateInput{
		Id:        req.Id,
		UpdatedBy: contexts.GetUserId(ctx),
		TenantCreateUpdateBase: dto.TenantCreateUpdateBase{
			TenantName: strings.TrimSpace(req.TenantName), // 租户名称
			//PackageId:    req.PackageId,
			LinkMan: strings.TrimSpace(req.LinkMan), // 联系人姓名
			//LinkPhone:    strings.TrimSpace(req.LinkPhone), // 联系人手机
			Status:       int(req.Status),                // 租户状态 1正常/2停用
			Website:      strings.TrimSpace(req.Website), // 绑定域名
			ExpireAt:     req.ExpireAt,                   // 过期时间
			AccountCount: req.AccountCount,               // 账号数量
		},
	})
}
