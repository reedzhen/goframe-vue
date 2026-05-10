package admin

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"goframe-vben/internal/library/gftenant"
	"goframe-vben/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"

	"goframe-vben/api/admin/common"
)

func (c *ControllerCommon) TenantGetInfoByHost(ctx context.Context, req *common.TenantGetInfoByHostReq) (res *common.TenantGetInfoByHostRes, err error) {
	// 非贴牌不返回数据
	mode := g.Cfg().MustGet(ctx, "tenant.mode").String()
	if mode == gftenant.ModeNone {
		return nil, nil
	}

	tenant, err := service.Tenant().GetInfoWithCache(g.RequestFromCtx(ctx))
	if err != nil || tenant == nil {
		return nil, gerror.New("公司不存在，请联系管理员")
	}

	return &common.TenantGetInfoByHostRes{
		Id:         tenant.Id,
		TenantName: tenant.TenantName,
		Website:    tenant.Website,
		ExpireAt:   tenant.ExpireAt,
	}, nil
}
