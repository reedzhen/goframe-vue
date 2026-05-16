// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-05-16 22:56:55
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysRoleDao is the data access object for the table sys_role.
type SysRoleDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SysRoleColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SysRoleColumns defines and stores column names for the table sys_role.
type SysRoleColumns struct {
	Id         string // 角色id
	ParentId   string // 上级角色ID
	Name       string // 角色名称
	Code       string // 角色标识
	Level      string // 关系树等级
	Tree       string // 关系树
	DataScope  string // 数据范围 1全部/2当前部门/3当前及以下部门/4自定义部门
	CustomDept string // 自定义部门权限
	Status     string // 角色状态 1正常/2停用
	Sort       string // 排序
	Remark     string // 备注
	CreatedBy  string // 添加人
	CreatedAt  string // 创建时间
	UpdatedBy  string // 更新人
	UpdatedAt  string // 更新时间
	TenantId   string // 租户ID
}

// sysRoleColumns holds the columns for the table sys_role.
var sysRoleColumns = SysRoleColumns{
	Id:         "id",
	ParentId:   "parent_id",
	Name:       "name",
	Code:       "code",
	Level:      "level",
	Tree:       "tree",
	DataScope:  "data_scope",
	CustomDept: "custom_dept",
	Status:     "status",
	Sort:       "sort",
	Remark:     "remark",
	CreatedBy:  "created_by",
	CreatedAt:  "created_at",
	UpdatedBy:  "updated_by",
	UpdatedAt:  "updated_at",
	TenantId:   "tenant_id",
}

// NewSysRoleDao creates and returns a new DAO object for table data access.
func NewSysRoleDao(handlers ...gdb.ModelHandler) *SysRoleDao {
	return &SysRoleDao{
		group:    "default",
		table:    "sys_role",
		columns:  sysRoleColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysRoleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysRoleDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysRoleDao) Columns() SysRoleColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysRoleDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysRoleDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysRoleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
