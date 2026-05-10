package consts

// 贴牌模式缓存Key
const (
	CacheTenantDomainKey    = "tenant:domain:%s"        // 租户通过域名缓存 SelectCache:tenant_cln.frp.deepic.cn
	CacheTenantIdKey        = "tenant:%d"               // 租户通过租户编号缓存
	CacheTenantConfigKey    = "cache:tenant_%d:conf_%s" // 租户配置缓存
	CacheTenantConfigKeyTTL = 12 * 60 * 60              // 租户配置缓存时间（s）
)

//// 角色接口权限缓存Key
//const (
//	CacheRoleApiPathKey    = "cache:api_path:role_%d" // 角色接口权限缓存
//	CacheRoleApiPathKeyTTL = 7 * 24 * 60 * 60         // 角色接口权限缓存时间（s）
//)

// 非贴牌模式缓存Key
const (
	CacheConfigKey    = "cache:config:group_%s" // 系统配置缓存
	CacheConfigKeyTTL = 12 * 60 * 60            // 系统配置缓存时间（s）
)
