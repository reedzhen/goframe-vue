package plugin

import (
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"goframe-vben/internal/library/addons"
	"goframe-vben/internal/library/asynqueue"
	"goframe-vben/internal/task"
)

type Asynq struct {
}

func (p *Asynq) Name() string {
	return "asynq"
}

func (p *Asynq) Author() string {
	return "Larry Liu"
}

func (p *Asynq) Version() string {
	return "1.0"
}

func (p *Asynq) Description() string {
	return "queue service"
}

func (p *Asynq) Install(s *ghttp.Server) error {
	ctx := gctx.GetInitCtx()

	// 获取 redis 配置
	var cfg gredis.Config
	if err := g.Cfg().MustGet(ctx, "redis.default").Scan(&cfg); err != nil {
		return err
	}

	// 初始化客户端
	asynqueue.NewClient(cfg)
	// 初始化服务端
	asynqueue.NewServer(cfg)

	// 注册主模块处理方法
	asynqueue.Server.RegisterHandle(task.TypePing, task.HandlePingTask)
	// 注册插件处理方法
	addons.RegisterJobs(ctx)

	// 启动服务端服务
	return asynqueue.Server.Start()
}

func (p *Asynq) Remove() error {
	asynqueue.Server.Shutdown()
	return nil
}
