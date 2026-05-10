// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-03-29 15:49:32
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysLoginRecord is the golang structure of table sys_login_record for DAO operations like Where/Data.
type SysLoginRecord struct {
	g.Meta    `orm:"table:sys_login_record, do:true"`
	Id        any         // 主键
	Username  any         // 用户账号
	Os        any         // 操作系统
	Device    any         // 设备名
	Browser   any         // 浏览器类型
	Ip        any         // ip地址
	IpCity    any         // ip归属地
	LoginType any         // 操作类型 1登录成功/2登录失败/3退出登录
	Remark    any         // 备注
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	TenantId  any         // 租户ID
}
