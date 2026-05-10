// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-03-29 15:49:32
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysCronRecordDao is the data access object for the table sys_cron_record.
type SysCronRecordDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  SysCronRecordColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// SysCronRecordColumns defines and stores column names for the table sys_cron_record.
type SysCronRecordColumns struct {
	Id        string //
	CronId    string // 任务Id sys_cron.id
	SpendTime string // 消耗时间, 单位毫秒
	Output    string // 执行结果或输出，可能包括错误信息
	Status    string // 执行状态 success成功/failure失败
	Remark    string // 备注
	CreatedAt string // 创建时间
}

// sysCronRecordColumns holds the columns for the table sys_cron_record.
var sysCronRecordColumns = SysCronRecordColumns{
	Id:        "id",
	CronId:    "cron_id",
	SpendTime: "spend_time",
	Output:    "output",
	Status:    "status",
	Remark:    "remark",
	CreatedAt: "created_at",
}

// NewSysCronRecordDao creates and returns a new DAO object for table data access.
func NewSysCronRecordDao(handlers ...gdb.ModelHandler) *SysCronRecordDao {
	return &SysCronRecordDao{
		group:    "default",
		table:    "sys_cron_record",
		columns:  sysCronRecordColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysCronRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysCronRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysCronRecordDao) Columns() SysCronRecordColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysCronRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysCronRecordDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysCronRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
