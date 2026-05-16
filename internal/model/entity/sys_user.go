// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-05-16 22:56:55
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUser is the golang structure for table sys_user.
type SysUser struct {
	Id             int64       `json:"id"             orm:"id"              ` //
	OrganizationId int64       `json:"organizationId" orm:"organization_id" ` // 机构Id
	Nickname       string      `json:"nickname"       orm:"nickname"        ` // 昵称
	Username       string      `json:"username"       orm:"username"        ` // 账号
	Password       string      `json:"password"       orm:"password"        ` // 密码
	Salt           string      `json:"salt"           orm:"salt"            ` // 盐加密
	RoleId         int64       `json:"roleId"         orm:"role_id"         ` // 角色Id
	Phone          string      `json:"phone"          orm:"phone"           ` // 手机号
	Avatar         string      `json:"avatar"         orm:"avatar"          ` // 头像
	Email          string      `json:"email"          orm:"email"           ` // 邮箱
	EmailVerified  uint        `json:"emailVerified"  orm:"email_verified"  ` // 邮箱是否验证:  1是 2否
	RealName       string      `json:"realName"       orm:"real_name"       ` // 真实姓名
	IdCard         string      `json:"idCard"         orm:"id_card"         ` // 身份证号
	Birthday       *gtime.Time `json:"birthday"       orm:"birthday"        ` // 出生日期
	Introduction   string      `json:"introduction"   orm:"introduction"    ` // 个人简介
	Status         int         `json:"status"         orm:"status"          ` // 状态 1正常/2冻结
	LastLoginAt    *gtime.Time `json:"lastLoginAt"    orm:"last_login_at"   ` // 最新一次登录时间
	ParentId       int64       `json:"parentId"       orm:"parent_id"       ` // 上级用户Id
	Level          int         `json:"level"          orm:"level"           ` // 关系树等级
	Tree           string      `json:"tree"           orm:"tree"            ` // 关系树
	IsAdmin        int         `json:"isAdmin"        orm:"is_admin"        ` // 是否超级管理员 1是/2否 默认：2
	CreatedBy      int64       `json:"createdBy"      orm:"created_by"      ` // 添加人
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      ` // 创建时间
	UpdatedBy      int64       `json:"updatedBy"      orm:"updated_by"      ` // 更新人
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      ` // 更新时间
	TenantId       int64       `json:"tenantId"       orm:"tenant_id"       ` // 租户ID
}
