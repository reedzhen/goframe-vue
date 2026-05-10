package admin

import (
	"context"
	"goframe-vben/internal/library/contexts"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/role"
)

func (c *ControllerRole) Update(ctx context.Context, req *role.UpdateReq) (res *role.UpdateRes, err error) {
	return nil, service.Role().Update(ctx, dto.RoleUpdateInput{
		Id:     req.Id,
		Name:   req.Name,
		Code:   req.Code,
		Remark: req.Remark,
		//ParentId:  req.ParentId,
		Status:    int(req.Status),
		UpdatedBy: contexts.GetUserId(ctx),
		//Permissions: req.Permissions,
	})
}
