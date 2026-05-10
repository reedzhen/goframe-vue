package admin

import (
	"context"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/loginrecord"
)

func (c *ControllerLoginrecord) GetPage(ctx context.Context, req *loginrecord.GetPageReq) (res *loginrecord.GetPageRes, err error) {
	data, err := service.LoginRecord().Page(ctx, dto.LoginRecordPageInput{
		PageParam: req.PageParam,
		Username:  req.Username,
		CreatedAt: req.CreatedAt,
	})

	return &loginrecord.GetPageRes{Result: data}, err
}
