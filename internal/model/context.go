package model

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Context 请求上下文结构
type Context struct {
	Module    string    // 应用模块 admin/api/open
	AddonName string    // 插件名称 如果不是插件模块请求，可能为空
	User      *Identity // 上下文用户信息
	Data      g.Map     // 自定KV变量，业务模块根据需要设置，不固定
}

// Identity 请求上下文中的用户信息
type Identity struct {
	UserId         int64  // 用户Id
	Username       string // 用户账号
	Nickname       string // 用户昵称
	OrganizationId int64  // 组织Id
	IsAdmin        bool   // 是否超管
	RoleId         int64  // 角色Id
}
