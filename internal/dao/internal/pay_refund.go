// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-05-16 22:56:55
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PayRefundDao is the data access object for the table pay_refund.
type PayRefundDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  PayRefundColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// PayRefundColumns defines and stores column names for the table pay_refund.
type PayRefundColumns struct {
	Id                string //
	No                string // 退款单号
	AppId             string // 应用ID
	ChannelId         string // 渠道ID
	ChannelCode       string // 渠道编码 例如:alipay_pc/wx_lite
	PayOrderId        string // 支付订单ID
	PayOrderNo        string // 支付订单编号
	MerchantOrderId   string // 商户订单ID 例：商城订单ID
	MerchantRefundId  string // 商户退款订单ID
	NotifyUrl         string // 异步通知商户地址
	Status            string // 退款状态 0未退款/10退款成功/20退款失败
	PayPrice          string // 支付金额 单位分
	RefundPrice       string // 退款金额 单位分
	Reason            string // 申请退款原因
	UserIp            string // 用户IP
	ChannelOrderNo    string // 渠道订单号，pay_order 中的 channel_order_no 对应
	ChannelRefundNo   string // 渠道退款单号，渠道返回
	SuccessTime       string // 退款成功时间
	ChannelErrorCode  string // 渠道调用报错时，错误码
	ChannelErrorMsg   string // 渠道调用报错时，错误信息
	ChannelNotifyData string // 支付渠道异步通知的内容
	Remark            string // 退款备注
	CreatedBy         string // 添加人
	CreatedAt         string // 创建时间
	UpdatedBy         string // 更新人
	UpdatedAt         string // 更新时间
	TenantId          string // 租户ID
}

// payRefundColumns holds the columns for the table pay_refund.
var payRefundColumns = PayRefundColumns{
	Id:                "id",
	No:                "no",
	AppId:             "app_id",
	ChannelId:         "channel_id",
	ChannelCode:       "channel_code",
	PayOrderId:        "pay_order_id",
	PayOrderNo:        "pay_order_no",
	MerchantOrderId:   "merchant_order_id",
	MerchantRefundId:  "merchant_refund_id",
	NotifyUrl:         "notify_url",
	Status:            "status",
	PayPrice:          "pay_price",
	RefundPrice:       "refund_price",
	Reason:            "reason",
	UserIp:            "user_ip",
	ChannelOrderNo:    "channel_order_no",
	ChannelRefundNo:   "channel_refund_no",
	SuccessTime:       "success_time",
	ChannelErrorCode:  "channel_error_code",
	ChannelErrorMsg:   "channel_error_msg",
	ChannelNotifyData: "channel_notify_data",
	Remark:            "remark",
	CreatedBy:         "created_by",
	CreatedAt:         "created_at",
	UpdatedBy:         "updated_by",
	UpdatedAt:         "updated_at",
	TenantId:          "tenant_id",
}

// NewPayRefundDao creates and returns a new DAO object for table data access.
func NewPayRefundDao(handlers ...gdb.ModelHandler) *PayRefundDao {
	return &PayRefundDao{
		group:    "default",
		table:    "pay_refund",
		columns:  payRefundColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PayRefundDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PayRefundDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PayRefundDao) Columns() PayRefundColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PayRefundDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PayRefundDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PayRefundDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
