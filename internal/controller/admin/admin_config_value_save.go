package admin

import (
	"context"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/config"
)

func (c *ControllerConfig) ValueSave(ctx context.Context, req *config.ValueSaveReq) (res *config.ValueSaveRes, err error) {
	values := make([]dto.ConfigValueItemInput, 0, len(req.Values))
	for _, v := range req.Values {
		values = append(values, dto.ConfigValueItemInput{
			ConfigKey:   v.ConfigKey,
			ConfigValue: v.ConfigValue,
		})
	}
	return nil, service.Config().SaveValuesByModuleCode(ctx, dto.ConfigSaveValuesInput{
		ModuleCode: req.ModuleCode,
		Values:     values,
	})
}
