// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-03-29 15:49:32
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PayOrderExtensionDao is the data access object for the table pay_order_extension.
type PayOrderExtensionDao struct {
	table    string                   // table is the underlying table name of the DAO.
	group    string                   // group is the database configuration group name of the current DAO.
	columns  PayOrderExtensionColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler       // handlers for customized model modification.
}

// PayOrderExtensionColumns defines and stores column names for the table pay_order_extension.
type PayOrderExtensionColumns struct {
	Id                string //
	No                string // 支付订单号
	PayOrderId        string // 支付订单编号
	ChannelId         string // 渠道编号
	ChannelCode       string // 渠道编码
	UserIp            string // 用户 IP
	Status            string // 支付状态 1未支付/10支付成功/20已退款/30支付关闭
	ChannelExtras     string // 支付渠道的额外参数
	ChannelErrorCode  string // 渠道调用报错时，错误码
	ChannelErrorMsg   string // 渠道调用报错时，错误信息
	ChannelNotifyData string // 支付渠道异步通知的内容
	CreatedBy         string // 添加人
	CreatedAt         string // 创建时间
	UpdatedBy         string // 更新人
	UpdatedAt         string // 更新时间
	TenantId          string // 租户ID
}

// payOrderExtensionColumns holds the columns for the table pay_order_extension.
var payOrderExtensionColumns = PayOrderExtensionColumns{
	Id:                "id",
	No:                "no",
	PayOrderId:        "pay_order_id",
	ChannelId:         "channel_id",
	ChannelCode:       "channel_code",
	UserIp:            "user_ip",
	Status:            "status",
	ChannelExtras:     "channel_extras",
	ChannelErrorCode:  "channel_error_code",
	ChannelErrorMsg:   "channel_error_msg",
	ChannelNotifyData: "channel_notify_data",
	CreatedBy:         "created_by",
	CreatedAt:         "created_at",
	UpdatedBy:         "updated_by",
	UpdatedAt:         "updated_at",
	TenantId:          "tenant_id",
}

// NewPayOrderExtensionDao creates and returns a new DAO object for table data access.
func NewPayOrderExtensionDao(handlers ...gdb.ModelHandler) *PayOrderExtensionDao {
	return &PayOrderExtensionDao{
		group:    "default",
		table:    "pay_order_extension",
		columns:  payOrderExtensionColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PayOrderExtensionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PayOrderExtensionDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PayOrderExtensionDao) Columns() PayOrderExtensionColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PayOrderExtensionDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PayOrderExtensionDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PayOrderExtensionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
