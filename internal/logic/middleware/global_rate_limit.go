package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"goframe-vben/internal/codes"
	"goframe-vben/internal/library/ratelimit"
	"goframe-vben/utility/response"
)

// GlobalRateLimit 全局 IP 限流中间件
// 大厂标准做法：所有接口都要有基础限流保护
// 默认：200 次/分钟/IP
func (s *sMiddleware) GlobalRateLimit(r *ghttp.Request) {
	ctx := r.GetCtx()
	ip := r.GetClientIp()

	// 执行限流检查（200次/分钟）
	if err := ratelimit.CheckIPLimit(ctx, 200); err != nil {
		g.Log().Warningf(ctx, "全局IP限流: %s, 路径: %s", ip, r.URL.Path)
		response.JsonExit(r, codes.CodeTooManyRequests.Code(), "请求过于频繁，请稍后再试")
		return
	}

	// 继续处理请求
	r.Middleware.Next()
}
