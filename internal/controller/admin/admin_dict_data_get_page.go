package admin

import (
	"context"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/dict"
)

func (c *ControllerDict) DataGetPage(ctx context.Context, req *dict.DataGetPageReq) (res *dict.DataGetPageRes, err error) {
	data, err := service.DictData().Page(ctx, dto.DictDataPageInput{
		PageParam: req.PageParam,
		DictId:    req.DictId,
		Keywords:  req.Keywords,
	})
	return &dict.DataGetPageRes{Result: data}, err
}
