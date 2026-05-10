// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-03-29 15:49:32
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysTenantPackage is the golang structure of table sys_tenant_package for DAO operations like Where/Data.
type SysTenantPackage struct {
	g.Meta    `orm:"table:sys_tenant_package, do:true"`
	Id        any         // 套餐编号
	Name      any         // 套餐名
	Status    any         // 套餐状态 1正常/2停用
	MenuIds   any         // 菜单ID 逗号分隔
	Remark    any         // 备注
	CreatedBy any         // 添加人
	CreatedAt *gtime.Time // 创建时间
	UpdatedBy any         // 更新人
	UpdatedAt *gtime.Time // 更新时间
}
