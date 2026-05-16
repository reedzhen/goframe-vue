// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-05-16 22:56:55
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysConfig is the golang structure for table sys_config.
type SysConfig struct {
	Id           int64       `json:"id"           orm:"id"            ` //
	Group        string      `json:"group"        orm:"group"         ` // 配置分组
	Label        string      `json:"label"        orm:"label"         ` // 中文名称 例如:姓名
	Key          string      `json:"key"          orm:"key"           ` // 配置键名 例如:age
	Value        string      `json:"value"        orm:"value"         ` // 配置值 例如:18
	DefaultValue string      `json:"defaultValue" orm:"default_value" ` // 默认值
	Type         string      `json:"type"         orm:"type"          ` // 类型 string,int,uint,bool,datetime,date
	Desc         string      `json:"desc"         orm:"desc"          ` // 描述
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    ` // 创建时间
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    ` // 更新时间
	TenantId     int64       `json:"tenantId"     orm:"tenant_id"     ` // 贴牌Id
}
