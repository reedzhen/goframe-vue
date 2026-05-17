package admin

import (
	"context"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/config"
)

func (c *ControllerConfig) ModuleDelete(ctx context.Context, req *config.ModuleDeleteReq) (res *config.ModuleDeleteRes, err error) {
	return nil, service.Config().ModuleDelete(ctx, req.Id)
}
