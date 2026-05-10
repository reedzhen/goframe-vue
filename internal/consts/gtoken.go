package consts

const (
	TokenType            = "Bearer"
	MultiLogin           = true
	FrontendMultiLogin   = true
	GTokenExpireIn       = 10 * 24 * 60 * 60 * 1000
	GTokenAdminPrefix    = "Admin:" // gtoken登录 管理后台 前缀区分
	GTokenFrontendPrefix = "Api:"   // gtoken登录 前台用户 前缀区分
	GTokeOpenPrefix      = "Open:"  // gtoken登录 第三方用户 前缀区分
)
