package middleware

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-vben/internal/library/cache"
	"goframe-vben/utility/response"
	"goframe-vben/utility/tools"
)

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
func (s *sMiddleware) CheckSum(r *ghttp.Request) {
	ctx := r.GetCtx()

	// 获取请求头中的签名参数
	// accessKey := r.Request.Header.Get("AccessKey") // 可选：用于区分不同第三方
	timestamp := r.Request.Header.Get("Timestamp")
	nonce := r.Request.Header.Get("NonceStr")
	checkSum := r.Request.Header.Get("CheckSum")

	// 参数校验
	if timestamp == "" || nonce == "" || checkSum == "" {
		response.JsonExit(r, gcode.CodeValidationFailed.Code(), "缺少必要的请求头: Timestamp, NonceStr, CheckSum")
		return
	}

	// 通过 AccessKey 获取密钥（实际项目应该从数据库或配置中心获取）
	accessKeySecret := "secret123"

	// 1. 时间戳校验：请求不能超过5分钟
	ts := gconv.Int64(timestamp)
	if gtime.Now().Unix()-ts > 300 {
		g.Log().Warningf(ctx, "时间戳过期: timestamp=%s, client_ip=%s", timestamp, r.GetClientIp())
		response.JsonExit(r, gcode.CodeValidationFailed.Code(), "时间戳过期")
		return
	}

	// 2. Nonce 防重放校验（防止请求被重复提交）
	if err := s.checkNonceForCheckSum(ctx, nonce); err != nil {
		g.Log().Warningf(ctx, "Nonce重复使用: nonce=%s, client_ip=%s", nonce, r.GetClientIp())
		response.JsonExit(r, gcode.CodeValidationFailed.Code(), "请求重复，请勿重复提交")
		return
	}

	// 3. 签名校验：验证请求完整性
	if !tools.ValidateCheckSum(checkSum, timestamp, nonce, accessKeySecret) {
		g.Log().Warningf(ctx, "签名错误: client_ip=%s, path=%s", r.GetClientIp(), r.URL.Path)
		response.JsonExit(r, gcode.CodeValidationFailed.Code(), "签名错误")
		return
	}

	// 所有校验通过，继续处理请求
	r.Middleware.Next()
}

// checkNonceForCheckSum 检查 Nonce 是否已被使用（防重放）
func (s *sMiddleware) checkNonceForCheckSum(ctx context.Context, nonce string) error {
	cacheInstance := cache.Instance()

	// Nonce 缓存键
	nonceKey := fmt.Sprintf("checksum:nonce:%s", nonce)

	// 检查 Nonce 是否已存在
	exists, err := cacheInstance.Contains(ctx, nonceKey)
	if err != nil {
		// 缓存查询失败，记录日志但不阻断（避免影响正常业务）
		g.Log().Errorf(ctx, "查询Nonce缓存失败: %v", err)
		return nil
	}

	if exists {
		return gerror.New("Nonce已被使用")
	}

	// 记录 Nonce，过期时间设置为 6 分钟（比时间窗口多1分钟）
	// 这样可以确保在时间窗口内的 Nonce 不会被重复使用
	if err = cacheInstance.Set(ctx, nonceKey, 1, 360); err != nil {
		g.Log().Errorf(ctx, "设置Nonce缓存失败: %v", err)
		// 缓存设置失败不阻断请求
		return nil
	}

	return nil
}

// 如何测试
// 1. 以下代码放入apifox或者postman的接口前置操作中
// 2. 然后Headers里面引入下面js代码里设置的环境变量
//
// 引入 crypto-js
// const CryptoJS = require('crypto-js');
// 获取当前时间戳
// var timestamp = Math.round(new Date().getTime() / 1000); // 注意：时间戳通常是秒级的
// 随机字符串
// var nonceStr = pm.variables.replaceIn('{{$guid}}');
// Access Key 和 Secret
// var accessKey = "key123";
// var secret = "secret123";
// 构建签名字符串
// var sign = secret + nonceStr + timestamp;
// 计算 CheckSum
// var checkSum = CryptoJS.SHA1(sign).toString(CryptoJS.enc.Hex).toLowerCase();
// 设置环境变量
// pm.environment.set("Timestamp", timestamp);
// pm.environment.set("NonceStr", nonceStr);
// pm.environment.set("AccessKey", accessKey);
// pm.environment.set("CheckSum", checkSum);
