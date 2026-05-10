// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-03-29 15:49:32
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUser is the golang structure of table sys_user for DAO operations like Where/Data.
type SysUser struct {
	g.Meta         `orm:"table:sys_user, do:true"`
	Id             any         //
	OrganizationId any         // 机构Id
	Nickname       any         // 昵称
	Username       any         // 账号
	Password       any         // 密码
	Salt           any         // 盐加密
	RoleId         any         // 角色Id
	Phone          any         // 手机号
	Avatar         any         // 头像
	Email          any         // 邮箱
	EmailVerified  any         // 邮箱是否验证:  1是 2否
	RealName       any         // 真实姓名
	IdCard         any         // 身份证号
	Birthday       *gtime.Time // 出生日期
	Introduction   any         // 个人简介
	Status         any         // 状态 1正常/2冻结
	LastLoginAt    *gtime.Time // 最新一次登录时间
	ParentId       any         // 上级用户Id
	Level          any         // 关系树等级
	Tree           any         // 关系树
	IsAdmin        any         // 是否超级管理员 1是/2否 默认：2
	CreatedBy      any         // 添加人
	CreatedAt      *gtime.Time // 创建时间
	UpdatedBy      any         // 更新人
	UpdatedAt      *gtime.Time // 更新时间
	TenantId       any         // 租户ID
}
