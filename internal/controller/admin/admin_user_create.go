package admin

import (
	"context"
	"goframe-vben/internal/library/contexts"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"
	"strings"

	"goframe-vben/api/admin/user"
)

func (c *ControllerUser) Create(ctx context.Context, req *user.CreateReq) (res *user.CreateRes, err error) {
	_, err = service.User().Create(ctx, dto.UserCreateInput{
		ParentId:  contexts.GetUserId(ctx),
		Password:  req.Password,
		CreatedBy: contexts.GetUserId(ctx),
		UserCreateUpdateBase: dto.UserCreateUpdateBase{
			Username:       strings.TrimSpace(req.Username),
			Nickname:       strings.TrimSpace(req.Nickname),
			Phone:          strings.TrimSpace(req.Phone),
			RoleId:         req.RoleId,
			OrganizationId: req.OrganizationId,
		},
	})
	return nil, err
}
