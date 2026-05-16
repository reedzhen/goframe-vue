// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-05-16 22:56:55
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDict is the golang structure of table sys_dict for DAO operations like Where/Data.
type SysDict struct {
	g.Meta    `orm:"table:sys_dict, do:true"`
	Id        any         // 字典id
	Code      any         // 字典标识
	Name      any         // 字典名称
	Sort      any         // 排序号
	Note      any         // 备注
	CreatedBy any         // 添加人
	CreatedAt *gtime.Time // 创建时间
	UpdatedBy any         // 更新人
	UpdatedAt *gtime.Time // 更新时间
}
