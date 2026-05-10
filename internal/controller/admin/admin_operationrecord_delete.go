package admin

import (
	"context"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/operationrecord"
)

func (c *ControllerOperationrecord) Delete(ctx context.Context, req *operationrecord.DeleteReq) (res *operationrecord.DeleteRes, err error) {
	return nil, service.OperationRecord().Delete(ctx, req.Day)
}
