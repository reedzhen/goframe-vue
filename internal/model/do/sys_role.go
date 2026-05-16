// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-05-16 22:56:55
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRole is the golang structure of table sys_role for DAO operations like Where/Data.
type SysRole struct {
	g.Meta     `orm:"table:sys_role, do:true"`
	Id         any         // 角色id
	ParentId   any         // 上级角色ID
	Name       any         // 角色名称
	Code       any         // 角色标识
	Level      any         // 关系树等级
	Tree       any         // 关系树
	DataScope  any         // 数据范围 1全部/2当前部门/3当前及以下部门/4自定义部门
	CustomDept *gjson.Json // 自定义部门权限
	Status     any         // 角色状态 1正常/2停用
	Sort       any         // 排序
	Remark     any         // 备注
	CreatedBy  any         // 添加人
	CreatedAt  *gtime.Time // 创建时间
	UpdatedBy  any         // 更新人
	UpdatedAt  *gtime.Time // 更新时间
	TenantId   any         // 租户ID
}
