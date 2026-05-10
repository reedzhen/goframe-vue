// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IMiddleware interface {
		// CheckApiAuth 校验接口权限中间件
		CheckApiAuth(r *ghttp.Request)
		// CheckJwtToken token校验中间件
		CheckJwtToken(r *ghttp.Request)
		// CheckSum 接口签名校验中间件（适用于第三方开放平台）
		// 功能：
		// 1. Timestamp 校验：防止旧请求重放（时间窗口5分钟）
		// 2. Nonce 防重放：唯一请求标识，防止重复提交
		// 3. CheckSum 签名：验证请求未被篡改
		//
		// 使用场景：
		// - 第三方系统调用开放接口获取 Token
		// - 使用 Token 调用业务接口
		// - 类似微信/支付宝开放平台的签名验证
		CheckSum(r *ghttp.Request)
		// CORS 跨域
		CORS(r *ghttp.Request)
		// Ctx 请求结束时不会自动调用Done方法结束掉ctx
		Ctx(r *ghttp.Request)
		// GlobalRateLimit 全局 IP 限流中间件
		// 大厂标准做法：所有接口都要有基础限流保护
		// 默认：200 次/分钟/IP
		GlobalRateLimit(r *ghttp.Request)
		// HandlerResponse 自定义统一响应返回中间件
		HandlerResponse(r *ghttp.Request)
		// OperationRecord 记录操作日志
		OperationRecord(r *ghttp.Request)
		// ResponseHandler 返回处理中间件
		ResponseHandler(r *ghttp.Request)
		// SetTenant 校验和设置租户
		SetTenant(r *ghttp.Request)
	}
)

var (
	localMiddleware IMiddleware
)

func Middleware() IMiddleware {
	if localMiddleware == nil {
		panic("implement not found for interface IMiddleware, forgot register?")
	}
	return localMiddleware
}

func RegisterMiddleware(i IMiddleware) {
	localMiddleware = i
}
