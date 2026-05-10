package boot

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gmode"
	"goframe-vben/internal/library/cache"
	"goframe-vben/internal/library/migrate"
	"strings"
)

func Init(ctx context.Context) {
	// 全局设置i18n
	g.I18n().SetLanguage("zh-CN")

	// 检查数据库连接是否正常
	if err := ConnDb(); err != nil {
		g.Log().Fatal(ctx, "数据库连接失败，请检查数据库配置和网络连接", err)
	}

	// 执行数据库迁移，保证服务启动前表结构已更新
	if err := migrate.Run(ctx); err != nil {
		g.Log().Fatal(ctx, "数据库迁移失败", err)
	}

	// 设置当前程序运行环境
	SetGfMode(ctx)

	// 设置时区
	if err := SetTimeZone(); err != nil {
		g.Log().Fatal(ctx, "时区设置失败", err)
	}

	// 设置缓存适配器
	cache.SetAdapter()
}

// GetCurrentModule 获取当前接口请求前缀 admin-api/app-api/open-api
func GetCurrentModule(r *ghttp.Request) string {
	slice := strings.Split(r.URL.Path, "/")
	if len(slice) < 2 || slice[1] == "" {
		return ""
	}

	return slice[1]
}

// SetGfMode 设置当前程序运行环境
func SetGfMode(ctx context.Context) {
	mode := g.Cfg().MustGet(ctx, "app.mode").String()
	if mode == "" || !gstr.InArray([]string{gmode.DEVELOP, gmode.TESTING, gmode.STAGING, gmode.PRODUCT}, mode) {
		mode = gmode.NOT_SET
	}

	gmode.Set(mode)
}

// SetTimeZone 设置时区
func SetTimeZone() error {
	return gtime.SetTimeZone("Asia/Shanghai")
}

// ConnDb 检查数据库连接是否正常
func ConnDb() error {
	return g.DB().PingMaster()
}
