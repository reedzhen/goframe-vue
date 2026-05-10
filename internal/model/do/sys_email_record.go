// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-03-29 15:49:32
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysEmailRecord is the golang structure of table sys_email_record for DAO operations like Where/Data.
type SysEmailRecord struct {
	g.Meta       `orm:"table:sys_email_record, do:true"`
	Id           any         // 主键id
	Title        any         // 邮件标题
	Content      any         // 邮件内容
	Receiver     any         // 收件邮箱
	Sender       any         // 发件邮箱
	CreateUserId any         // 创建人
	Note         any         // 备注
	TenantId     any         // 租户id
	CreateTime   *gtime.Time // 创建时间
	UpdateTime   *gtime.Time // 修改时间
}
