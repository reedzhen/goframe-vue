// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-05-16 22:56:55
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysDictDao is the data access object for the table sys_dict.
type SysDictDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SysDictColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SysDictColumns defines and stores column names for the table sys_dict.
type SysDictColumns struct {
	Id        string // 字典id
	Code      string // 字典标识
	Name      string // 字典名称
	Sort      string // 排序号
	Note      string // 备注
	CreatedBy string // 添加人
	CreatedAt string // 创建时间
	UpdatedBy string // 更新人
	UpdatedAt string // 更新时间
}

// sysDictColumns holds the columns for the table sys_dict.
var sysDictColumns = SysDictColumns{
	Id:        "id",
	Code:      "code",
	Name:      "name",
	Sort:      "sort",
	Note:      "note",
	CreatedBy: "created_by",
	CreatedAt: "created_at",
	UpdatedBy: "updated_by",
	UpdatedAt: "updated_at",
}

// NewSysDictDao creates and returns a new DAO object for table data access.
func NewSysDictDao(handlers ...gdb.ModelHandler) *SysDictDao {
	return &SysDictDao{
		group:    "default",
		table:    "sys_dict",
		columns:  sysDictColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysDictDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysDictDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysDictDao) Columns() SysDictColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysDictDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysDictDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysDictDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
