package admin

import (
	"context"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/role"
)

func (c *ControllerRole) SaveMenuList(ctx context.Context, req *role.SaveMenuListReq) (res *role.SaveMenuListRes, err error) {
	return nil, service.Role().SaveRoleMenuList(ctx, dto.RoleSaveMenuInput{
		RoleId:  req.Id,
		MenuIds: req.MenuIds,
	})
}
