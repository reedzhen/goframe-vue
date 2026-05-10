package admin

import (
	"context"
	"goframe-vben/internal/library/contexts"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/role"
)

func (c *ControllerRole) Create(ctx context.Context, req *role.CreateReq) (res *role.CreateRes, err error) {
	_, err = service.Role().Create(ctx, dto.RoleCreateInput{
		Name:      req.Name,
		Code:      req.Code,
		Remark:    req.Remark,
		Status:    int(req.Status),
		CreatedBy: contexts.GetUserId(ctx),
	})
	return nil, err
}
