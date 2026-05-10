package admin

import (
	"context"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/dict"
)

func (c *ControllerDict) Update(ctx context.Context, req *dict.UpdateReq) (res *dict.UpdateRes, err error) {
	return nil, service.Dict().Update(ctx, dto.DictUpdateInput{
		Id:   req.Id,
		Code: req.Code,
		Name: req.Name,
		Note: req.Note,
	})
}
