// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-03-29 15:49:32
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysCronDao is the data access object for the table sys_cron.
type SysCronDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SysCronColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SysCronColumns defines and stores column names for the table sys_cron.
type SysCronColumns struct {
	Id        string // 任务ID
	Title     string // 标题
	Tag       string // 标签 字典表获取
	ApiUrl    string // 接口地址
	ApiParam  string // 接口请求参数
	ApiHeader string // 接口请求头
	Pattern   string // cron表达式
	Status    string // 任务状态 active运行中/inactive已结束
	Policy    string // 策略 1并行/2单例/3单次/4多次
	Count     string // 执行次数 policy=4时有效
	Sort      string // 排序
	NextRunAt string // 下次预计运行时间
	LastRunAt string // 最后一次运行时间
	Remark    string // 备注
	CreatedBy string // 添加人
	CreatedAt string // 创建时间
	UpdatedBy string // 更新人
	UpdatedAt string // 更新时间
}

// sysCronColumns holds the columns for the table sys_cron.
var sysCronColumns = SysCronColumns{
	Id:        "id",
	Title:     "title",
	Tag:       "tag",
	ApiUrl:    "api_url",
	ApiParam:  "api_param",
	ApiHeader: "api_header",
	Pattern:   "pattern",
	Status:    "status",
	Policy:    "policy",
	Count:     "count",
	Sort:      "sort",
	NextRunAt: "next_run_at",
	LastRunAt: "last_run_at",
	Remark:    "remark",
	CreatedBy: "created_by",
	CreatedAt: "created_at",
	UpdatedBy: "updated_by",
	UpdatedAt: "updated_at",
}

// NewSysCronDao creates and returns a new DAO object for table data access.
func NewSysCronDao(handlers ...gdb.ModelHandler) *SysCronDao {
	return &SysCronDao{
		group:    "default",
		table:    "sys_cron",
		columns:  sysCronColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysCronDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysCronDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysCronDao) Columns() SysCronColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysCronDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysCronDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysCronDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
