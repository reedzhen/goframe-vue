// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-05-16 22:56:55
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysConfigItem is the golang structure for table sys_config_item.
type SysConfigItem struct {
	Id           int64       `json:"id"           orm:"id"            ` // 主键ID
	ModuleId     int64       `json:"moduleId"     orm:"module_id"     ` // 所属配置模块ID
	Name         string      `json:"name"         orm:"name"          ` // 配置项名称，后台展示用
	ConfigKey    string      `json:"configKey"    orm:"config_key"    ` // 配置项键名，代码读取用唯一标识
	ConfigValue  string      `json:"configValue"  orm:"config_value"  ` // 配置项当前值
	DefaultValue string      `json:"defaultValue" orm:"default_value" ` // 默认值，读取失败或未配置时兜底
	ValueType    string      `json:"valueType"    orm:"value_type"    ` // 值类型 string/int/uint/bool/json/datetime/date
	InputType    int         `json:"inputType"    orm:"input_type"    ` // 输入类型 1输入框/2范围/3下拉/4单选/5开关/6多选
	InputParams  string      `json:"inputParams"  orm:"input_params"  ` // 控件参数，选项或范围配置，例如 A-1|B-2|C-3
	Description  string      `json:"description"  orm:"description"   ` // 配置项说明
	Sort         int         `json:"sort"         orm:"sort"          ` // 排序值，越小越靠前
	Status       int         `json:"status"       orm:"status"        ` // 状态 1启用/0禁用
	IsSystem     int         `json:"isSystem"     orm:"is_system"     ` // 是否系统内置 1是/0否，内置配置通常由研发维护
	CreatedBy    int64       `json:"createdBy"    orm:"created_by"    ` // 添加人
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    ` // 创建时间
	UpdatedBy    int64       `json:"updatedBy"    orm:"updated_by"    ` // 更新人
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    ` // 更新时间
	TenantId     int64       `json:"tenantId"     orm:"tenant_id"     ` // 租户ID
}
