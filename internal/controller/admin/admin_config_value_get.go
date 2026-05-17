package admin

import (
	"context"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/config"
)

func (c *ControllerConfig) ValueGet(ctx context.Context, req *config.ValueGetReq) (res *config.ValueGetRes, err error) {
	data, err := service.Config().GetValuesByModuleCode(ctx, req.ModuleCode)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, nil
	}
	return &config.ValueGetRes{
		ModuleCode: data.ModuleCode,
		Data:       data.Data,
	}, nil
}
