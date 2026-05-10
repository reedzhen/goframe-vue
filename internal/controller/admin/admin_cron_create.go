package admin

import (
	"context"
	"goframe-vben/internal/consts"
	"goframe-vben/internal/library/contexts"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"
	"strings"

	"goframe-vben/api/admin/cron"
)

func (c *ControllerCron) Create(ctx context.Context, req *cron.CreateReq) (res *cron.CreateRes, err error) {
	return nil, service.Cron().Create(ctx, dto.CronCreateInput{
		CreatedBy: contexts.GetUserId(ctx),
		Status:    consts.CronStatusInactive, // 任务状态 1运行中/2已结束
		CronCreateUpdateBase: dto.CronCreateUpdateBase{
			Title:     strings.TrimSpace(req.Title),   // 标题
			Tag:       req.Tag,                        // 标签 字典表获取
			ApiUrl:    strings.TrimSpace(req.ApiUrl),  // 接口地址
			ApiParam:  req.ApiParam,                   // 接口请求参数
			ApiHeader: req.ApiHeader,                  // 接口请求头
			Pattern:   strings.TrimSpace(req.Pattern), // cron表达式
			Policy:    req.Policy,                     // 策略 1并行/2单例/3单次/4多次
			Count:     req.Count,
			Sort:      req.Sort,                      // 排序
			Remark:    strings.TrimSpace(req.Remark), // 备注
		},
	})
}
