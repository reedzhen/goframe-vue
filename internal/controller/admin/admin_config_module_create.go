package admin

import (
	"context"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/config"
)

func (c *ControllerConfig) ModuleCreate(ctx context.Context, req *config.ModuleCreateReq) (res *config.ModuleCreateRes, err error) {
	return nil, service.Config().ModuleCreate(ctx, dto.ConfigModuleCreateInput{
		Code:        req.Code,
		Name:        req.Name,
		Description: req.Description,
		Sort:        req.Sort,
		Status:      req.Status,
	})
}
