package admin

import (
	"context"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/role"
)

func (c *ControllerRole) GetList(ctx context.Context, req *role.GetListReq) (res *role.GetListRes, err error) {
	list, err := service.Role().GetList(ctx, dto.RoleGetListInput{
		Name: req.Name,
	})
	if err != nil {
		return
	}

	return &role.GetListRes{List: list}, nil
}
