package admin

import (
	"context"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/cron"
)

func (c *ControllerCron) ExecuteOnce(ctx context.Context, req *cron.ExecuteOnceReq) (res *cron.ExecuteOnceRes, err error) {
	return nil, service.Cron().ExecuteOnce(ctx, req.Id)
}
