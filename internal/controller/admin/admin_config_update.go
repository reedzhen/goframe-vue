package admin

import (
	"context"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/config"
)

func (c *ControllerConfig) Update(ctx context.Context, req *config.UpdateReq) (res *config.UpdateRes, err error) {
	return nil, service.Config().UpdateByGroup(ctx, dto.ConfigUpdateInput{
		Group: req.Group,
		Data:  req.Data,
	})
}
