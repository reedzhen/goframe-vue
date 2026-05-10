// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-03-29 15:49:32
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysConfig is the golang structure of table sys_config for DAO operations like Where/Data.
type SysConfig struct {
	g.Meta       `orm:"table:sys_config, do:true"`
	Id           any         //
	Group        any         // 配置分组
	Label        any         // 中文名称 例如:姓名
	Key          any         // 配置键名 例如:age
	Value        any         // 配置值 例如:18
	DefaultValue any         // 默认值
	Type         any         // 类型 string,int,uint,bool,datetime,date
	Desc         any         // 描述
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 更新时间
	TenantId     any         // 贴牌Id
}
