package admin

import (
	"context"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/operationrecord"
)

func (c *ControllerOperationrecord) GetPage(ctx context.Context, req *operationrecord.GetPageReq) (res *operationrecord.GetPageRes, err error) {
	data, err := service.OperationRecord().Page(ctx, dto.OperationRecordPageInput{
		PageParam: req.PageParam,
		Username:  req.Username,
		Module:    req.Module,
		CreatedAt: req.CreatedAt,
	})

	return (*operationrecord.GetPageRes)(data), err
}
