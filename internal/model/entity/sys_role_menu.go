// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-03-29 15:49:32
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRoleMenu is the golang structure for table sys_role_menu.
type SysRoleMenu struct {
	Id        int64       `json:"id"        orm:"id"         ` // 主键id
	RoleId    int64       `json:"roleId"    orm:"role_id"    ` // 角色id
	MenuId    int64       `json:"menuId"    orm:"menu_id"    ` // 菜单id
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
	TenantId  int64       `json:"tenantId"  orm:"tenant_id"  ` // 租户ID
}
