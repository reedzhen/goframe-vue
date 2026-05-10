package admin

import (
	"context"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/role"
)

func (c *ControllerRole) GetMenuList(ctx context.Context, req *role.GetMenuListReq) (res *role.GetMenuListRes, err error) {
	list, err := service.Role().GetRoleMenuList(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return (*role.GetMenuListRes)(&list), nil
}
