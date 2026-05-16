// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-05-16 22:56:55
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRoleMenu is the golang structure of table sys_role_menu for DAO operations like Where/Data.
type SysRoleMenu struct {
	g.Meta    `orm:"table:sys_role_menu, do:true"`
	Id        any         // 主键id
	RoleId    any         // 角色id
	MenuId    any         // 菜单id
	CreatedAt *gtime.Time // 创建时间
	TenantId  any         // 租户ID
}
