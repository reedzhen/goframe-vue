package dto

import (
	"github.com/gogf/gf/v2/database/gdb"
	"goframe-vben/internal/library/query"
	"goframe-vben/internal/model/entity"
	"strings"
)

// UserLoginInput 用户登录
type UserLoginInput struct {
	LoginType string // account phone
	Username  string // 账号
	Password  string // 密码(明文)
}

// UserPageInput 用户分页
type UserPageInput struct {
	query.PageParam
	Nickname       string
	Username       string
	OrganizationId int64 // 组织Id
}

func (q *UserPageInput) Cond(m *gdb.Model) *gdb.Model {
	//ctx := m.GetCtx()
	if q.Nickname != "" {
		m = m.WhereLike("sys_user.nickname", "%"+strings.TrimSpace(q.Nickname)+"%")
	}
	if q.Username != "" {
		m = m.WhereLike("sys_user.username", "%"+strings.TrimSpace(q.Username)+"%")
	}
	if q.OrganizationId != 0 {
		m = m.Where("sys_user.organization_id", q.OrganizationId)
	}

	//// 非超管用户只能查看自己和自己的子级（当然也看不到超管）
	//// 超管看不到自己
	//if contexts.GetUserId(ctx) != consts.SuperUserId {
	//	u := contexts.GetUser(ctx)
	//	m = m.Where("sys_user.id = ? OR sys_user.tree like ?", u.UserId, tools.TrimLikeRight(tree.BuildChildPath(u.Tree, u.UserId)))
	//	// 只查看当前组织下的用户
	//	m = m.Where("sys_user.org_id", contexts.GetOrgId(ctx))
	//} else {
	//	m = m.WhereNot("sys_user.id ", consts.SuperUserId)
	//}
	return m
}

// UserPageOutput 用户分页返回
type UserPageOutput struct {
	*entity.SysUser
	Role         *entity.SysRole         `json:"role" orm:"with:id=role_id"`
	Organization *entity.SysOrganization `json:"organization" orm:"with:id=organization_id"`
}

// UserCreateInput 用户新增
type UserCreateInput struct {
	ParentId  int64  // 上级用户ID
	Password  string // 密码
	TenantId  int64  // 租户ID
	CreatedBy int64
	UserCreateUpdateBase
}

type UserCreateUpdateBase struct {
	Username       string // 账号
	Nickname       string // 昵称
	Phone          string // 手机号
	RoleId         int64  // 角色ID
	OrganizationId int64  // 机构Id
}

// UserUpdateInput 用户编辑
type UserUpdateInput struct {
	Id int64 // 主键ID
	UserCreateUpdateBase
	UpdatedBy int64
}

// UserGetListInput 用户列表
type UserGetListInput struct {
	OrganizationId int64  // 机构Id
	Nickname       string // 昵称
	Username       string // 账号
}
