// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-05-16 22:56:55
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysRoleMenuDao is the data access object for the table sys_role_menu.
type SysRoleMenuDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SysRoleMenuColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SysRoleMenuColumns defines and stores column names for the table sys_role_menu.
type SysRoleMenuColumns struct {
	Id        string // 主键id
	RoleId    string // 角色id
	MenuId    string // 菜单id
	CreatedAt string // 创建时间
	TenantId  string // 租户ID
}

// sysRoleMenuColumns holds the columns for the table sys_role_menu.
var sysRoleMenuColumns = SysRoleMenuColumns{
	Id:        "id",
	RoleId:    "role_id",
	MenuId:    "menu_id",
	CreatedAt: "created_at",
	TenantId:  "tenant_id",
}

// NewSysRoleMenuDao creates and returns a new DAO object for table data access.
func NewSysRoleMenuDao(handlers ...gdb.ModelHandler) *SysRoleMenuDao {
	return &SysRoleMenuDao{
		group:    "default",
		table:    "sys_role_menu",
		columns:  sysRoleMenuColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysRoleMenuDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysRoleMenuDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysRoleMenuDao) Columns() SysRoleMenuColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysRoleMenuDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysRoleMenuDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysRoleMenuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
