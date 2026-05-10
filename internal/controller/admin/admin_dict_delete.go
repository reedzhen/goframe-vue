package admin

import (
	"context"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/dict"
)

func (c *ControllerDict) Delete(ctx context.Context, req *dict.DeleteReq) (res *dict.DeleteRes, err error) {
	return nil, service.Dict().Delete(ctx, req.Id)
}
