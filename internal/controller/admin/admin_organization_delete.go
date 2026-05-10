package admin

import (
	"context"
	"goframe-vben/api/admin/organization"
	"goframe-vben/internal/service"
)

// Delete 删除
func (c *ControllerOrganization) Delete(ctx context.Context, req *organization.DeleteReq) (res *organization.DeleteRes, err error) {
	return nil, service.Organization().Delete(ctx, req.Id)
}
