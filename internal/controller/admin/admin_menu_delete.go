package admin

import (
	"context"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/menu"
)

func (c *ControllerMenu) Delete(ctx context.Context, req *menu.DeleteReq) (res *menu.DeleteRes, err error) {
	return nil, service.Menu().Delete(ctx, req.Id)
}
