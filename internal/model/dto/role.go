package dto

import (
	"github.com/gogf/gf/v2/database/gdb"
	"goframe-vben/internal/library/query"
	"goframe-vben/internal/model/entity"
	"strings"
)

// RolePageInput 角色分页
type RolePageInput struct {
	query.PageParam
	Name string
	Code string
}

func (q *RolePageInput) Cond(m *gdb.Model) *gdb.Model {
	if q.Name != "" {
		m = m.WhereLike("sys_role.name", "%"+strings.TrimSpace(q.Name)+"%")
	}
	if q.Code != "" {
		m = m.WhereLike("sys_role.code", "%"+strings.TrimSpace(q.Code)+"%")
	}
	return m
}

// RolePageOutput 角色分页返回
type RolePageOutput struct {
	*entity.SysRole
	Permissions []int64 `json:"permissions"` // 菜单Id
}

// RoleCreateInput 角色新增
type RoleCreateInput struct {
	//TenantId int64  // 租户Id
	Name   string // 角色名称
	Code   string // 角色标识
	Remark string // 备注
	//ParentId  int64  // 上级角色ID
	Status int // 角色状态 1正常/2停用
	//Level     int    // 关系树等级
	//Tree      string // 关系树
	CreatedBy int64
	//Permissions []int64 // 菜单Id
}

// RoleUpdateInput 角色编辑
type RoleUpdateInput struct {
	Id     int64  // 角色id
	Name   string // 角色名称
	Code   string // 角色标识
	Remark string // 备注
	//ParentId  int64  // 上级角色ID
	Status    int // 角色状态 1正常/2停用
	UpdatedBy int64
	//Permissions []int64 // 菜单Id
}

// RoleGetListInput 角色列表
type RoleGetListInput struct {
	Name string
}

// RoleSaveMenuInput 保存角色下所有菜单
type RoleSaveMenuInput struct {
	//TenantId int64   // 租户Id
	RoleId  int64   // 角色id
	MenuIds []int64 // 待保存的菜单
}
