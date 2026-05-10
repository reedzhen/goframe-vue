// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-03-29 15:49:32
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysEmailRecord is the golang structure for table sys_email_record.
type SysEmailRecord struct {
	Id           int         `json:"id"           orm:"id"             ` // 主键id
	Title        string      `json:"title"        orm:"title"          ` // 邮件标题
	Content      string      `json:"content"      orm:"content"        ` // 邮件内容
	Receiver     string      `json:"receiver"     orm:"receiver"       ` // 收件邮箱
	Sender       string      `json:"sender"       orm:"sender"         ` // 发件邮箱
	CreateUserId int         `json:"createUserId" orm:"create_user_id" ` // 创建人
	Note         string      `json:"note"         orm:"note"           ` // 备注
	TenantId     int         `json:"tenantId"     orm:"tenant_id"      ` // 租户id
	CreateTime   *gtime.Time `json:"createTime"   orm:"create_time"    ` // 创建时间
	UpdateTime   *gtime.Time `json:"updateTime"   orm:"update_time"    ` // 修改时间
}
