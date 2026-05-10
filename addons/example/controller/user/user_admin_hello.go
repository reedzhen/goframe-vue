package user

import (
	"context"
	"goframe-vben/addons/example/service"

	"goframe-vben/addons/example/api/user/admin"
)

func (c *ControllerAdmin) Hello(ctx context.Context, req *admin.HelloReq) (res *admin.HelloRes, err error) {
	user, err := service.User().GetProfileById(ctx, 1)
	return &admin.HelloRes{SysUser: user}, err
}
