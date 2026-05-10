// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-03-29 15:49:32
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysLoginRecord is the golang structure for table sys_login_record.
type SysLoginRecord struct {
	Id        int64       `json:"id"        orm:"id"         ` // 主键
	Username  string      `json:"username"  orm:"username"   ` // 用户账号
	Os        string      `json:"os"        orm:"os"         ` // 操作系统
	Device    string      `json:"device"    orm:"device"     ` // 设备名
	Browser   string      `json:"browser"   orm:"browser"    ` // 浏览器类型
	Ip        string      `json:"ip"        orm:"ip"         ` // ip地址
	IpCity    string      `json:"ipCity"    orm:"ip_city"    ` // ip归属地
	LoginType uint        `json:"loginType" orm:"login_type" ` // 操作类型 1登录成功/2登录失败/3退出登录
	Remark    string      `json:"remark"    orm:"remark"     ` // 备注
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 更新时间
	TenantId  int64       `json:"tenantId"  orm:"tenant_id"  ` // 租户ID
}
