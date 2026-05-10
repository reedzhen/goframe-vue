package admin

import (
	"context"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/user"
)

func (c *ControllerUser) GetPage(ctx context.Context, req *user.GetPageReq) (res *user.GetPageRes, err error) {
	data, err := service.User().Page(ctx, dto.UserPageInput{
		PageParam:      req.PageParam,
		Nickname:       req.Nickname,
		Username:       req.Username,
		OrganizationId: req.OrganizationId,
	})
	if err != nil {
		return
	}

	return &user.GetPageRes{Result: data}, nil
}
