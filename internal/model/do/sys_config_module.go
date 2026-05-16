// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-05-16 22:56:55
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysConfigModule is the golang structure of table sys_config_module for DAO operations like Where/Data.
type SysConfigModule struct {
	g.Meta      `orm:"table:sys_config_module, do:true"`
	Id          any         // 主键ID
	Code        any         // 模块编码，代码中使用的唯一标识
	Name        any         // 模块名称
	Description any         // 模块说明
	Sort        any         // 排序值，越小越靠前
	Status      any         // 状态 1启用/0禁用
	CreatedBy   any         // 添加人
	CreatedAt   *gtime.Time // 创建时间
	UpdatedBy   any         // 更新人
	UpdatedAt   *gtime.Time // 更新时间
	TenantId    any         // 租户ID
}
