package tenant

import (
	"context"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"
	"strings"

	"goframe-vben/api/tenant/admin"
)

func (c *ControllerAdmin) GetList(ctx context.Context, req *admin.GetListReq) (res *admin.GetListRes, err error) {
	list, err := service.Tenant().GetList(ctx, dto.TenantGetListInput{
		TenantName: strings.TrimSpace(req.TenantName), // 租户名称
		LinkMan:    strings.TrimSpace(req.LinkMan),    // 联系人姓名
		LinkPhone:  strings.TrimSpace(req.LinkPhone),  // 联系人手机
	})
	if err != nil {
		return nil, err
	}

	return &admin.GetListRes{List: list}, nil
}
