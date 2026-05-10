package cache

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
)

// cache 缓存驱动
var cache *gcache.Cache

// Instance 缓存实例
func Instance() *gcache.Cache {
	if cache == nil {
		panic("缓存未初始化")
	}
	return cache
}

// SetAdapter 设置缓存适配器
func SetAdapter() {
	// 默认使用redis缓存
	adapter := gcache.NewAdapterRedis(g.Redis("default"))

	// db查询缓存
	g.DB().GetCache().SetAdapter(adapter)
	// 贴牌db模式数据库是动态切的，这里会存在很多贴牌，所以查询缓存使用cache.Instance().Set()
	//g.DB("tenant").GetCache().SetAdapter(adapter)

	// 通用缓存 cache.Instance().Set()
	cache = gcache.New()
	cache.SetAdapter(adapter)
}
