package admin

import (
	"context"
	"goframe-vben/internal/library/contexts"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/cron"
)

func (c *ControllerCron) Update(ctx context.Context, req *cron.UpdateReq) (res *cron.UpdateRes, err error) {
	return nil, service.Cron().Update(ctx, dto.CronUpdateInput{
		Id:        req.Id,
		UpdatedBy: contexts.GetUserId(ctx),
		CronCreateUpdateBase: dto.CronCreateUpdateBase{
			Title:     req.Title,     // 标题
			Tag:       req.Tag,       // 标签 字典表获取
			ApiUrl:    req.ApiUrl,    // 接口地址
			ApiParam:  req.ApiParam,  // 接口请求参数
			ApiHeader: req.ApiHeader, // 接口请求头
			Pattern:   req.Pattern,   // cron表达式
			Policy:    req.Policy,    // 策略 1并行/2单例/3单次/4多次
			Count:     req.Count,
			Sort:      req.Sort,   // 排序
			Remark:    req.Remark, // 备注
		},
	})
}
