package addons

import (
	"context"
	"sync"

	"github.com/gogf/gf/v2/net/ghttp"
)

// Addon 插件
type Addon interface {
	Name() string
	//Init(ctx context.Context)
	InitRouter(ctx context.Context, group *ghttp.RouterGroup)
	InitJob(ctx context.Context)
}

var (
	addons = make(map[string]Addon)
	mLock  sync.Mutex
)

// RegisterModulesRouter 注册所有已安装模块路由
func RegisterModulesRouter(ctx context.Context, group *ghttp.RouterGroup) {
	for _, module := range addons {
		module.InitRouter(ctx, group)
	}
}

// RegisterJobs 注册所有已安装模块脚本
func RegisterJobs(ctx context.Context) {
	for _, module := range addons {
		module.InitJob(ctx)
	}
}

// RegisterModule 注册模块
func RegisterModule(m Addon) Addon {
	mLock.Lock()
	defer mLock.Unlock()
	name := m.Name()
	if _, ok := addons[name]; ok {
		panic("module repeat registration, name:" + name)
	}
	addons[name] = m
	return m
}
