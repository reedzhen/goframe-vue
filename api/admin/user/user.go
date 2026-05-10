package user

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-vben/internal/consts"
	"goframe-vben/internal/library/query"
)

type GetPageReq struct {
	g.Meta `path:"/user/page" method:"get" tags:"用户管理" summary:"用户分页"`
	query.PageParam
	Nickname       string `json:"nickname" dc:"用户昵称"`
	Username       string `json:"username" dc:"登录账号"`
	OrganizationId int64  `json:"organizationId" dc:"组织Id"`
}
type GetPageRes struct {
	*query.Result
}

type CreateReq struct {
	g.Meta         `path:"/user/create" method:"post" tags:"用户管理" summary:"新增用户"`
	Username       string `json:"username" v:"required|passport" dc:"用户账号"`
	Phone          string `json:"phone"  v:"required|phone" dc:"手机号"`
	Password       string `json:"password" v:"required" dc:"密码"`
	Nickname       string `json:"nickname" v:"required" dc:"昵称"`
	RoleId         int64  `json:"roleId" v:"required" dc:"角色id"`
	OrganizationId int64  `json:"organizationId" v:"required" dc:"机构Id"`
}
type CreateRes struct{}

type UpdateReq struct {
	g.Meta         `path:"/user/update" method:"post" tags:"用户管理" summary:"编辑用户"`
	Id             int64  `json:"id" v:"required" dc:"用户Id"`
	Username       string `json:"username" v:"passport"  dc:"用户账号"  `
	Phone          string `json:"phone" v:"required" dc:"手机号"`
	Nickname       string `json:"nickname" v:"required" dc:"昵称" `
	RoleId         int64  `json:"roleId" v:"required" dc:"角色Id"`
	OrganizationId int64  `json:"organizationId" v:"required" dc:"机构Id"`
}
type UpdateRes struct{}

type DeleteReq struct {
	g.Meta `path:"/user/delete/{Id}" method:"post" tags:"用户管理" summary:"删除用户"`
	Id     int64 `json:"id" in:"path" v:"required"  dc:"用户Id"`
}
type DeleteRes struct{}

type CheckFieldExistReq struct {
	g.Meta `path:"/user/exist" method:"get" tags:"用户管理" summary:"校验字段唯一" noAuth:"true"`
	Field  string `json:"field" v:"required" in:"query" dc:"字段"`
	Value  string `json:"value" v:"required" in:"query" dc:"值"`
	Id     int64  `json:"id" in:"query"  dc:"用户id(不包含此用户id)"`
}
type CheckFieldExistRes struct{}

type UpdatePwdReq struct {
	g.Meta      `path:"/user/update-pwd" method:"post" tags:"用户管理" summary:"修改密码"`
	OldPassword string `json:"oldPassword" v:"required"`
	NewPassword string `json:"newPassword" v:"required|password"`
}
type UpdatePwdRes struct{}

type ResetPwdReq struct {
	g.Meta   `path:"/user/reset-pwd" method:"post" tags:"用户管理" summary:"重置密码"`
	UserId   int64  `json:"userId" v:"required"`
	Password string `json:"password" v:"required|password"`
}
type ResetPwdRes struct{}

type ChangeStatusReq struct {
	g.Meta `path:"/user/change-status" method:"post" tags:"用户管理" summary:"修改用户状态"`
	UserId int64             `json:"userId" v:"required" dc:"用户id"`
	Status consts.UserStatus `json:"status" v:"enums" dc:"1:可用;2:冻结"`
}
type ChangeStatusRes struct{}

type GetListReq struct {
	g.Meta         `path:"/user/list" method:"get" tags:"用户管理" summary:"获取用户列表"`
	OrganizationId int64  `json:"organizationId" v:"required" dc:"机构Id"`
	Nickname       string `json:"nickname" in:"query" dc:"昵称"`
	Username       string `json:"username" in:"query" dc:"账号"`
}
type GetListRes []*UserItem

// UserItem 用户表包含密码等敏感数据，这里只返回非敏感数据
type UserItem struct {
	Id       int64  `json:"id" dc:"用户id"`
	Username string `json:"username"  dc:"账号"`
	Nickname string `json:"nickname"  dc:"昵称"`
	RealName string `json:"realName"  dc:"昵称"`
	Avatar   string `json:"avatar"  dc:"头像"`
}
