package admin

import (
	"context"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/config"
)

func (c *ControllerConfig) ModuleUpdate(ctx context.Context, req *config.ModuleUpdateReq) (res *config.ModuleUpdateRes, err error) {
	return nil, service.Config().ModuleUpdate(ctx, dto.ConfigModuleUpdateInput{
		Id:          req.Id,
		Code:        req.Code,
		Name:        req.Name,
		Description: req.Description,
		Sort:        req.Sort,
		Status:      req.Status,
	})
}
