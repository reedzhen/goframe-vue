package admin

import (
	"context"
	"goframe-vben/internal/library/contexts"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/user"
)

func (c *ControllerUser) Update(ctx context.Context, req *user.UpdateReq) (res *user.UpdateRes, err error) {
	return nil, service.User().Update(ctx, dto.UserUpdateInput{
		Id:        req.Id,
		UpdatedBy: contexts.GetUserId(ctx),
		UserCreateUpdateBase: dto.UserCreateUpdateBase{
			Username:       req.Username,
			Nickname:       req.Nickname,
			Phone:          req.Phone,
			RoleId:         req.RoleId,
			OrganizationId: req.OrganizationId,
		},
	})
}
