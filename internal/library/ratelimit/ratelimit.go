package ratelimit

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"goframe-vben/internal/library/cache"
	"time"
)

// CheckIPLimit 检查 IP 限流
// 适用场景：公开接口的通用限流、全局 IP 限流
// 参数：
//   - maxRequests: 最大请求数（默认60，全局限流建议200）
func CheckIPLimit(ctx context.Context, maxRequests ...int) error {
	ip := g.RequestFromCtx(ctx).GetClientIp()

	limit := 60 // 默认 60 次/分钟
	if len(maxRequests) > 0 && maxRequests[0] > 0 {
		limit = maxRequests[0]
	}

	return CheckRateLimit(ctx, fmt.Sprintf("rate_limit:ip:%s", ip), limit, time.Minute)
}

// CheckUserLimit 检查用户限流
// 适用场景：已登录用户的接口限流
func CheckUserLimit(ctx context.Context, userId int64, maxRequests ...int) error {
	if userId == 0 {
		return CheckIPLimit(ctx, maxRequests...)
	}

	limit := 100 // 默认 100 次/分钟
	if len(maxRequests) > 0 && maxRequests[0] > 0 {
		limit = maxRequests[0]
	}

	return CheckRateLimit(ctx, fmt.Sprintf("rate_limit:user:%d", userId), limit, time.Minute)
}

// CheckLoginLimit 检查登录限流
// 适用场景：登录、注册等认证接口
func CheckLoginLimit(ctx context.Context, maxRequests ...int) error {
	ip := g.RequestFromCtx(ctx).GetClientIp()

	limit := 5 // 默认 5 次/分钟
	if len(maxRequests) > 0 && maxRequests[0] > 0 {
		limit = maxRequests[0]
	}

	return CheckRateLimit(ctx, fmt.Sprintf("rate_limit:login:%s", ip), limit, time.Minute)
}

// CheckSmsLimit 检查短信限流
// 适用场景：短信验证码发送接口
func CheckSmsLimit(ctx context.Context, phone string) error {
	if phone == "" {
		return nil
	}

	// 1分钟限制：1次
	if err := CheckRateLimit(ctx, fmt.Sprintf("rate_limit:sms:minute:%s", phone), 1, time.Minute); err != nil {
		return gerror.New("短信发送过于频繁，请1分钟后再试")
	}

	// 24小时限制：10次
	if err := CheckRateLimit(ctx, fmt.Sprintf("rate_limit:sms:day:%s", phone), 10, 24*time.Hour); err != nil {
		return gerror.New("今日短信发送次数已达上限")
	}

	return nil
}

// CheckRateLimit 检查限流（滑动窗口算法）
func CheckRateLimit(ctx context.Context, cacheKey string, maxRequests int, window time.Duration) error {
	cacheInstance := cache.Instance()

	now := time.Now().UnixMilli()
	windowStart := now - window.Milliseconds()

	// 获取当前窗口内的请求记录
	requests, err := cacheInstance.Get(ctx, cacheKey)
	if err != nil {
		// 缓存查询失败，记录日志但不阻断
		g.Log().Errorf(ctx, "获取限流缓存失败: %v", err)
		return nil
	}

	// 解析请求时间列表
	var timestamps []int64
	if !requests.IsNil() {
		if err = requests.Scan(&timestamps); err != nil {
			g.Log().Errorf(ctx, "解析限流数据失败: %v", err)
			return nil
		}
	}

	// 移除窗口外的旧记录
	validTimestamps := make([]int64, 0)
	for _, ts := range timestamps {
		if ts > windowStart {
			validTimestamps = append(validTimestamps, ts)
		}
	}

	// 检查是否超过限制
	if len(validTimestamps) >= maxRequests {
		return gerror.New("请求过于频繁，请稍后再试")
	}

	// 添加当前请求时间
	validTimestamps = append(validTimestamps, now)

	// 更新缓存
	if err = cacheInstance.Set(ctx, cacheKey, validTimestamps, window+time.Second); err != nil {
		g.Log().Errorf(ctx, "设置限流缓存失败: %v", err)
		return nil
	}

	return nil
}
