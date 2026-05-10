package admin

import (
	"context"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/role"
)

func (c *ControllerRole) GetPage(ctx context.Context, req *role.GetPageReq) (res *role.GetPageRes, err error) {
	data, err := service.Role().Page(ctx, dto.RolePageInput{
		PageParam: req.PageParam,
		Name:      req.Name,
		Code:      req.Code,
	})
	return &role.GetPageRes{Result: data}, err
}
