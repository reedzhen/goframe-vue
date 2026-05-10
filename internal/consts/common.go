package consts

const (
	ContextHTTPKey  = "HttpContextKey" // http上下文变量存储键名
	TenantHeaderKey = "X-TenantId"     // 租户请求头名称
)

// 业务日志分组
const (
	LoggerGroupCron  = "cron"  // 定时任务
	LoggerGroupAsynq = "asynq" // asynq
)

// 路由前缀
const (
	GroupAdminApi = "admin-api"
	GroupAppApi   = "app-api"
	GroupOpenApi  = "open-api"
)
