package admin

import (
	"context"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/dict"
)

func (c *ControllerDict) DataDelete(ctx context.Context, req *dict.DataDeleteReq) (res *dict.DataDeleteRes, err error) {
	return nil, service.DictData().Delete(ctx, req.Id)
}
