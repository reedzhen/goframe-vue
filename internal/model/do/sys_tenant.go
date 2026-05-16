// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-05-16 22:56:55
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysTenant is the golang structure of table sys_tenant for DAO operations like Where/Data.
type SysTenant struct {
	g.Meta       `orm:"table:sys_tenant, do:true"`
	Id           any         //
	TenantName   any         // 租户名称
	LinkId       any         // 联系人ID sys_user.id
	LinkMan      any         // 联系人姓名
	LinkPhone    any         // 联系人手机
	Status       any         // 租户状态 1正常/2停用
	Website      any         // 绑定域名
	PackageId    any         // 租户套餐ID
	ExpireAt     *gtime.Time // 过期时间
	AccountCount any         // 账号数量
	CreatedBy    any         // 添加人
	CreatedAt    *gtime.Time // 创建时间
	UpdatedBy    any         // 更新人
	UpdatedAt    *gtime.Time // 更新时间
}
