package admin

import (
	"context"
	"goframe-vben/api/admin/organization"
	"goframe-vben/internal/library/contexts"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"
	"strings"
)

// Update 修改
func (c *ControllerOrganization) Update(ctx context.Context, req *organization.UpdateReq) (res *organization.UpdateRes, err error) {
	return nil, service.Organization().Update(ctx, dto.OrganizationUpdateInput{
		Id:        req.Id,
		UpdatedBy: contexts.GetUserId(ctx),
		OrganizationCreateUpdateBase: dto.OrganizationCreateUpdateBase{
			ParentId: req.ParentId,                    // 上级id, 0是顶级
			Code:     strings.TrimSpace(req.Code),     // 机构代码
			Name:     strings.TrimSpace(req.Name),     // 机构名称
			FullName: strings.TrimSpace(req.FullName), // 机构全称
			Type:     req.Type,                        // 机构类型
			Status:   req.Status,                      // 部门状态 1正常/2停用
			//LinkId:    req.LinkId,                       // 负责人id sys_user.id
			//LinkMan:   strings.TrimSpace(req.LinkMan),   // 联系人
			//LinkPhone: strings.TrimSpace(req.LinkPhone), // 联系电话
			Remark: strings.TrimSpace(req.Remark), // 备注
		},
	})
}
