package admin

import (
	"context"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/config"
)

func (c *ControllerConfig) ItemCreate(ctx context.Context, req *config.ItemCreateReq) (res *config.ItemCreateRes, err error) {
	return nil, service.Config().ItemCreate(ctx, dto.ConfigItemCreateInput{
		ModuleId:     req.ModuleId,
		Name:         req.Name,
		ConfigKey:    req.ConfigKey,
		ConfigValue:  req.ConfigValue,
		DefaultValue: req.DefaultValue,
		ValueType:    req.ValueType,
		InputType:    req.InputType,
		InputParams:  req.InputParams,
		Description:  req.Description,
		Sort:         req.Sort,
		Status:       req.Status,
		IsSystem:     req.IsSystem,
	})
}
