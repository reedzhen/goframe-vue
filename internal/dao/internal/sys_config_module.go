// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-05-16 22:56:55
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysConfigModuleDao is the data access object for the table sys_config_module.
type SysConfigModuleDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  SysConfigModuleColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// SysConfigModuleColumns defines and stores column names for the table sys_config_module.
type SysConfigModuleColumns struct {
	Id          string // 主键ID
	Code        string // 模块编码，代码中使用的唯一标识
	Name        string // 模块名称
	Description string // 模块说明
	Sort        string // 排序值，越小越靠前
	Status      string // 状态 1启用/0禁用
	CreatedBy   string // 添加人
	CreatedAt   string // 创建时间
	UpdatedBy   string // 更新人
	UpdatedAt   string // 更新时间
	TenantId    string // 租户ID
}

// sysConfigModuleColumns holds the columns for the table sys_config_module.
var sysConfigModuleColumns = SysConfigModuleColumns{
	Id:          "id",
	Code:        "code",
	Name:        "name",
	Description: "description",
	Sort:        "sort",
	Status:      "status",
	CreatedBy:   "created_by",
	CreatedAt:   "created_at",
	UpdatedBy:   "updated_by",
	UpdatedAt:   "updated_at",
	TenantId:    "tenant_id",
}

// NewSysConfigModuleDao creates and returns a new DAO object for table data access.
func NewSysConfigModuleDao(handlers ...gdb.ModelHandler) *SysConfigModuleDao {
	return &SysConfigModuleDao{
		group:    "default",
		table:    "sys_config_module",
		columns:  sysConfigModuleColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysConfigModuleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysConfigModuleDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysConfigModuleDao) Columns() SysConfigModuleColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysConfigModuleDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysConfigModuleDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysConfigModuleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
