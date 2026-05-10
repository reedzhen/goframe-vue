package admin

import (
	"context"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/config"
)

func (c *ControllerConfig) GetListByGroup(ctx context.Context, req *config.GetListByGroupReq) (res *config.GetListByGroupRes, err error) {
	data, err := service.Config().GetConfigByGroup(ctx, req.Group)
	if err != nil {
		return nil, err
	}

	return &config.GetListByGroupRes{
		Group: data.Group,
		Data:  data.Data,
	}, nil
}
