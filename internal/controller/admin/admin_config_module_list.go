package admin

import (
	"context"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/config"
)

func (c *ControllerConfig) ModuleList(ctx context.Context, req *config.ModuleListReq) (res *config.ModuleListRes, err error) {
	list, err := service.Config().ModuleList(ctx, dto.ConfigModuleListInput{
		Keywords: req.Keywords,
		Status:   req.Status,
	})
	if err != nil {
		return nil, err
	}

	return (*config.ModuleListRes)(&list), nil
}
