// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-03-29 15:49:32
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysFileRecordDao is the data access object for the table sys_file_record.
type SysFileRecordDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  SysFileRecordColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// SysFileRecordColumns defines and stores column names for the table sys_file_record.
type SysFileRecordColumns struct {
	Id        string // 主键id
	Name      string // 文件名称
	Path      string // 文件存储路径
	Length    string // 文件大小
	Type      string // 文件类型：img图片/file文件
	Note      string // 备注
	CreatedBy string // 创建人
	CreatedAt string // 创建时间
	UpdatedAt string // 修改时间
}

// sysFileRecordColumns holds the columns for the table sys_file_record.
var sysFileRecordColumns = SysFileRecordColumns{
	Id:        "id",
	Name:      "name",
	Path:      "path",
	Length:    "length",
	Type:      "type",
	Note:      "note",
	CreatedBy: "created_by",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewSysFileRecordDao creates and returns a new DAO object for table data access.
func NewSysFileRecordDao(handlers ...gdb.ModelHandler) *SysFileRecordDao {
	return &SysFileRecordDao{
		group:    "default",
		table:    "sys_file_record",
		columns:  sysFileRecordColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysFileRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysFileRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysFileRecordDao) Columns() SysFileRecordColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysFileRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysFileRecordDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysFileRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
