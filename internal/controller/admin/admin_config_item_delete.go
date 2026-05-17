package admin

import (
	"context"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/config"
)

func (c *ControllerConfig) ItemDelete(ctx context.Context, req *config.ItemDeleteReq) (res *config.ItemDeleteRes, err error) {
	return nil, service.Config().ItemDelete(ctx, req.Id)
}
