// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-03-29 15:49:32
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOrganization is the golang structure of table sys_organization for DAO operations like Where/Data.
type SysOrganization struct {
	g.Meta    `orm:"table:sys_organization, do:true"`
	Id        any         //
	ParentId  any         // 上级id, 0是顶级
	Code      any         // 机构代码
	Name      any         // 机构名称
	FullName  any         // 机构全称
	Type      any         // 机构类型
	Status    any         // 部门状态 1正常/2停用
	LinkId    any         // 负责人id sys_user.id
	LinkMan   any         // 联系人
	LinkPhone any         // 联系电话
	Level     any         // 关系树层级
	Tree      any         // 关系树
	Remark    any         // 备注
	CreatedBy any         // 添加人
	CreatedAt *gtime.Time // 创建时间
	UpdatedBy any         // 更新人
	UpdatedAt *gtime.Time // 更新时间
	TenantId  any         // 租户ID
}
