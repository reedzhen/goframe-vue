package admin

import (
	"context"
	"goframe-vben/api/admin/cron"
	"goframe-vben/internal/library/contexts"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"
)

func (c *ControllerCron) ChangeStatus(ctx context.Context, req *cron.ChangeStatusReq) (res *cron.ChangeStatusRes, err error) {
	err = service.Cron().ChangeStatus(ctx, dto.CronChangeStatusInput{
		Id:        req.Id,
		Status:    req.Status,
		UpdatedBy: contexts.GetUserId(ctx),
	})
	return nil, err
}
