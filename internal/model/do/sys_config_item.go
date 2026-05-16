// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-05-16 22:56:55
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysConfigItem is the golang structure of table sys_config_item for DAO operations like Where/Data.
type SysConfigItem struct {
	g.Meta       `orm:"table:sys_config_item, do:true"`
	Id           any         // 主键ID
	ModuleId     any         // 所属配置模块ID
	Name         any         // 配置项名称，后台展示用
	ConfigKey    any         // 配置项键名，代码读取用唯一标识
	ConfigValue  any         // 配置项当前值
	DefaultValue any         // 默认值，读取失败或未配置时兜底
	ValueType    any         // 值类型 string/int/uint/bool/json/datetime/date
	InputType    any         // 输入类型 1输入框/2范围/3下拉/4单选/5开关/6多选
	InputParams  any         // 控件参数，选项或范围配置，例如 A-1|B-2|C-3
	Description  any         // 配置项说明
	Sort         any         // 排序值，越小越靠前
	Status       any         // 状态 1启用/0禁用
	IsSystem     any         // 是否系统内置 1是/0否，内置配置通常由研发维护
	CreatedBy    any         // 添加人
	CreatedAt    *gtime.Time // 创建时间
	UpdatedBy    any         // 更新人
	UpdatedAt    *gtime.Time // 更新时间
	TenantId     any         // 租户ID
}
