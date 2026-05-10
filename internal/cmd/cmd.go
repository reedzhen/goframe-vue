package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
	"goframe-vben/internal/cmd/plugin"
)

func init() {
	if err := Main.AddCommand(&Code); err != nil {
		g.Log().Fatal(gctx.GetInitCtx(), "注册命令失败", err)
	}
}

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "启动所有服务",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// new(plugin.Cron)
			p := []ghttp.Plugin{new(plugin.Http), new(plugin.Asynq)}
			s.Plugin(p...)
			s.Run()

			return nil
		},
	}
)
