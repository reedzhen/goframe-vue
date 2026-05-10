package admin

import (
	"context"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/dict"
)

func (c *ControllerDict) Create(ctx context.Context, req *dict.CreateReq) (res *dict.CreateRes, err error) {
	return nil, service.Dict().Create(ctx, dto.DictCreateInput{
		Code: req.Code,
		Name: req.Name,
		Note: req.Note,
	})
}
