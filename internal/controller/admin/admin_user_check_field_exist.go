package admin

import (
	"context"
	"goframe-vben/api/admin/user"
	"goframe-vben/internal/service"
	"strings"
)

func (c *ControllerUser) CheckFieldExist(ctx context.Context, req *user.CheckFieldExistReq) (res *user.CheckFieldExistRes, err error) {
	return nil, service.User().CheckFieldExist(ctx, strings.TrimSpace(req.Field), strings.TrimSpace(req.Value), req.Id)
}
