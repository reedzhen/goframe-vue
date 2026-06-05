package admin

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	"goframe-vben/api/admin/user"
	"goframe-vben/internal/service"
)

func (c *ControllerUser) GetInfo(ctx context.Context, req *user.GetInfoReq) (res *user.GetInfoRes, err error) {
	info, err := service.User().ValidateExists(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if err = gconv.Struct(info, &res); err != nil {
		return nil, err
	}
	return res, nil
}
