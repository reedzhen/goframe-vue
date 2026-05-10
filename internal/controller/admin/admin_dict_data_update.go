package admin

import (
	"context"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/dict"
)

func (c *ControllerDict) DataUpdate(ctx context.Context, req *dict.DataUpdateReq) (res *dict.DataUpdateRes, err error) {
	return nil, service.DictData().Update(ctx, dto.DictDataUpdateInput{
		Id:     req.Id,
		DictId: req.DictId,
		Code:   req.Code,
		Name:   req.Name,
		Sort:   req.Sort,
		Note:   req.Note,
	})
}
