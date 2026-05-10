// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-03-29 15:49:32
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOrganization is the golang structure for table sys_organization.
type SysOrganization struct {
	Id        int64       `json:"id"        orm:"id"         ` //
	ParentId  int64       `json:"parentId"  orm:"parent_id"  ` // 上级id, 0是顶级
	Code      string      `json:"code"      orm:"code"       ` // 机构代码
	Name      string      `json:"name"      orm:"name"       ` // 机构名称
	FullName  string      `json:"fullName"  orm:"full_name"  ` // 机构全称
	Type      int         `json:"type"      orm:"type"       ` // 机构类型
	Status    int         `json:"status"    orm:"status"     ` // 部门状态 1正常/2停用
	LinkId    int64       `json:"linkId"    orm:"link_id"    ` // 负责人id sys_user.id
	LinkMan   string      `json:"linkMan"   orm:"link_man"   ` // 联系人
	LinkPhone string      `json:"linkPhone" orm:"link_phone" ` // 联系电话
	Level     int         `json:"level"     orm:"level"      ` // 关系树层级
	Tree      string      `json:"tree"      orm:"tree"       ` // 关系树
	Remark    string      `json:"remark"    orm:"remark"     ` // 备注
	CreatedBy int64       `json:"createdBy" orm:"created_by" ` // 添加人
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
	UpdatedBy int64       `json:"updatedBy" orm:"updated_by" ` // 更新人
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 更新时间
	TenantId  int64       `json:"tenantId"  orm:"tenant_id"  ` // 租户ID
}
