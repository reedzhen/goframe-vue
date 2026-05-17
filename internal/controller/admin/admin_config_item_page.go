package admin

import (
	"context"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/config"
)

func (c *ControllerConfig) ItemPage(ctx context.Context, req *config.ItemPageReq) (res *config.ItemPageRes, err error) {
	data, err := service.Config().ItemPage(ctx, dto.ConfigItemPageInput{
		PageParam: req.PageParam,
		ModuleId:  req.ModuleId,
		Keywords:  req.Keywords,
		Status:    req.Status,
	})
	if err != nil {
		return nil, err
	}
	return &config.ItemPageRes{Result: data}, nil
}
