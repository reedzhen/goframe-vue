// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-05-16 22:56:55
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysTenant is the golang structure for table sys_tenant.
type SysTenant struct {
	Id           int64       `json:"id"           orm:"id"            ` //
	TenantName   string      `json:"tenantName"   orm:"tenant_name"   ` // 租户名称
	LinkId       int64       `json:"linkId"       orm:"link_id"       ` // 联系人ID sys_user.id
	LinkMan      string      `json:"linkMan"      orm:"link_man"      ` // 联系人姓名
	LinkPhone    string      `json:"linkPhone"    orm:"link_phone"    ` // 联系人手机
	Status       int         `json:"status"       orm:"status"        ` // 租户状态 1正常/2停用
	Website      string      `json:"website"      orm:"website"       ` // 绑定域名
	PackageId    int64       `json:"packageId"    orm:"package_id"    ` // 租户套餐ID
	ExpireAt     *gtime.Time `json:"expireAt"     orm:"expire_at"     ` // 过期时间
	AccountCount int         `json:"accountCount" orm:"account_count" ` // 账号数量
	CreatedBy    int64       `json:"createdBy"    orm:"created_by"    ` // 添加人
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    ` // 创建时间
	UpdatedBy    int64       `json:"updatedBy"    orm:"updated_by"    ` // 更新人
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    ` // 更新时间
}
