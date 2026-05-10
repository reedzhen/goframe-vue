package admin

import (
	"context"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/dict"
)

func (c *ControllerDict) GetPage(ctx context.Context, req *dict.GetPageReq) (res *dict.GetPageRes, err error) {
	data, err := service.Dict().Page(ctx, dto.DictPageInput{
		PageParam: req.PageParam,
	})
	return &dict.GetPageRes{Result: data}, err
}
