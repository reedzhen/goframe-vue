package admin

import (
	"context"
	"goframe-vben/api/admin/organization"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"
	"strings"
)

// GetList 列表
func (c *ControllerOrganization) GetList(ctx context.Context, req *organization.GetListReq) (res *organization.GetListRes, err error) {
	list, err := service.Organization().GetList(ctx, dto.OrganizationGetListInput{
		Name: strings.TrimSpace(req.Name), // 机构名称
	})
	if err != nil {
		return nil, err
	}

	return (*organization.GetListRes)(&list), nil
}
