// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-05-16 22:56:55
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PayOrderDao is the data access object for the table pay_order.
type PayOrderDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  PayOrderColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// PayOrderColumns defines and stores column names for the table pay_order.
type PayOrderColumns struct {
	Id              string //
	AppId           string // 应用ID
	ChannelId       string // 渠道ID
	ChannelCode     string // 渠道编码 例如:alipay_pc/wx_lite
	MerchantOrderId string // 商户订单ID 例如：trade_order.id
	OrderGroup      string // 组别 mall商城
	Subject         string // 商品标题
	Body            string // 商品描述
	NotifyUrl       string // 支付通知回调地址
	PayType         string // 支付类型 wxpay微信支付/alipay支付宝/saobei扫呗
	TradeType       string // 交易类型 saobei(mini小程序) wxpay(mp公众号/mini小程序/appAPP/scan二维码扫码) alipay(web网页/appAPP/scan二维码扫码)
	Price           string // 支付金额 单位:分
	ChannelFeeRate  string // 渠道手续费率 单位:百分比
	ChannelFeePrice string // 渠道手续费金额 单位:分
	Status          string // 支付状态 1未支付/10支付成功/20已退款/30支付关闭
	UserIp          string // 用户IP
	ExpireTime      string // 订单失效时间
	SuccessTime     string // 订单支付成功时间
	ExtensionId     string // 订单拓展单编号 支付成功后写入
	No              string // 支付订单号 extension.no 支付成功后写入
	RefundPrice     string // 退款总金额 单位：分
	ChannelUserId   string // 渠道用户编号
	ChannelOrderNo  string // 渠道订单号
	CreatedBy       string // 添加人
	CreatedAt       string // 创建时间
	UpdatedBy       string // 更新人
	UpdatedAt       string // 更新时间
	TenantId        string // 租户ID
}

// payOrderColumns holds the columns for the table pay_order.
var payOrderColumns = PayOrderColumns{
	Id:              "id",
	AppId:           "app_id",
	ChannelId:       "channel_id",
	ChannelCode:     "channel_code",
	MerchantOrderId: "merchant_order_id",
	OrderGroup:      "order_group",
	Subject:         "subject",
	Body:            "body",
	NotifyUrl:       "notify_url",
	PayType:         "pay_type",
	TradeType:       "trade_type",
	Price:           "price",
	ChannelFeeRate:  "channel_fee_rate",
	ChannelFeePrice: "channel_fee_price",
	Status:          "status",
	UserIp:          "user_ip",
	ExpireTime:      "expire_time",
	SuccessTime:     "success_time",
	ExtensionId:     "extension_id",
	No:              "no",
	RefundPrice:     "refund_price",
	ChannelUserId:   "channel_user_id",
	ChannelOrderNo:  "channel_order_no",
	CreatedBy:       "created_by",
	CreatedAt:       "created_at",
	UpdatedBy:       "updated_by",
	UpdatedAt:       "updated_at",
	TenantId:        "tenant_id",
}

// NewPayOrderDao creates and returns a new DAO object for table data access.
func NewPayOrderDao(handlers ...gdb.ModelHandler) *PayOrderDao {
	return &PayOrderDao{
		group:    "default",
		table:    "pay_order",
		columns:  payOrderColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PayOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PayOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PayOrderDao) Columns() PayOrderColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PayOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PayOrderDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PayOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
