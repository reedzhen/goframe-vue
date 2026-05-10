package example

import (
	_ "goframe-vben/addons/example/logic"
	"goframe-vben/addons/example/router"
	"goframe-vben/addons/example/task"
	"goframe-vben/internal/library/addons"
	"goframe-vben/internal/library/asynqueue"

	"context"
	"github.com/gogf/gf/v2/net/ghttp"
)

func init() {
	addons.RegisterModule(&addon{})
}

type addon struct{}

// Name 插件名称
func (a *addon) Name() string {
	return "example"
}

func (a *addon) InitJob(ctx context.Context) {
	asynqueue.Server.RegisterHandle(task.TypeHello, task.HandleHelloTask)
}

func (a *addon) InitRouter(ctx context.Context, group *ghttp.RouterGroup) {
	router.Admin(ctx, group, a.Name())
}
