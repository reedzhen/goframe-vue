package admin

import (
	"context"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/config"
)

func (c *ControllerConfig) ItemUpdate(ctx context.Context, req *config.ItemUpdateReq) (res *config.ItemUpdateRes, err error) {
	return nil, service.Config().ItemUpdate(ctx, dto.ConfigItemUpdateInput{
		Id:           req.Id,
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
