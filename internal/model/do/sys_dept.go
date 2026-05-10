// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-03-29 15:49:32
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDept is the golang structure of table sys_dept for DAO operations like Where/Data.
type SysDept struct {
	g.Meta    `orm:"table:sys_dept, do:true"`
	Id        any         //
	ParentId  any         // 上级Id 0是顶级
	DeptCode  any         // 部门编码
	DeptName  any         // 部门名称
	Status    any         // 部门状态 1正常/2停用
	Level     any         // 关系树层级
	Tree      any         // 关系树
	Sort      any         // 排序
	Remark    any         // 备注
	CreatedBy any         // 添加人
	CreatedAt *gtime.Time // 创建时间
	UpdatedBy any         // 更新人
	UpdatedAt *gtime.Time // 更新时间
	TenantId  any         // 租户ID
}
