package admin

import (
	"context"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/dict"
)

func (c *ControllerDict) DataCreate(ctx context.Context, req *dict.DataCreateReq) (res *dict.DataCreateRes, err error) {
	return nil, service.DictData().Create(ctx, dto.DictDataCreateInput{
		DictId: req.DictId,
		Code:   req.Code,
		Name:   req.Name,
		Sort:   req.Sort,
		Note:   req.Note,
	})
}
