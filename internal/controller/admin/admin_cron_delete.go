package admin

import (
	"context"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/cron"
)

func (c *ControllerCron) Delete(ctx context.Context, req *cron.DeleteReq) (res *cron.DeleteRes, err error) {
	return nil, service.Cron().Delete(ctx, req.Id)
}
