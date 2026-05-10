package admin

import (
	"context"
	"goframe-vben/api/admin/cron"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"
)

func (c *ControllerCron) RecordPage(ctx context.Context, req *cron.RecordPageReq) (res *cron.RecordPageRes, err error) {
	data, err := service.CronRecord().Page(ctx, dto.CronRecordPageInput{
		PageParam: req.PageParam,
		CronId:    req.CronId,
	})
	return &cron.RecordPageRes{Result: data}, err
}
