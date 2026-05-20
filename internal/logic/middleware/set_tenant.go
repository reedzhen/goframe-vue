package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

// SetTenant 校验和设置租户
func (s *sMiddleware) SetTenant(r *ghttp.Request) {
	//ctx := r.GetCtx()

	//// 非贴牌直接过
	//mode := g.Cfg().MustGet(ctx, "tenant.mode").String()
	//if mode == gftenant.ModeNone {
	//	r.Middleware.Next()
	//	return
	//}
	//
	//// 列模式白名单接口
	//if mode == gftenant.ModeColumn {
	//	url := r.Request.URL.Path
	//	ignoreUrls := g.Cfg().MustGet(ctx, "tenant.ignoreUrls").Strings()
	//	for _, v := range ignoreUrls {
	//		if tools.KeyMatch(url, v) {
	//			r.Middleware.Next()
	//			return
	//		}
	//	}
	//}
	//
	//// 校验租户
	//tenant, err := service.Tenant().GetInfoWithCache(r)
	//if err != nil {
	//	response.JsonExit(r, codes.CodeBadRequest.Code(), err.Error())
	//}
	//if tenant == nil {
	//	response.JsonExit(r, codes.CodeBadRequest.Code(), "公司不存在，请联系管理员")
	//}
	//if tenant.Status != int(consts.TenantStatusOk) {
	//	response.JsonExit(r, codes.CodeBadRequest.Code(), "当前公司不可用，请联系管理员")
	//}
	//if gtime.Now().Timestamp() > tenant.ExpireAt.EndOfDay().Timestamp() {
	//	response.JsonExit(r, codes.CodeBadRequest.Code(), "当前授权已过期，请联系管理员")
	//}
	//
	//// 设置当前上下文中的租户ID
	//gftenant.SetTenant(ctx, tenant.Id)
	//
	//// db模式设置当前动态切换的数据库
	//if mode == gftenant.ModeDB {
	//	gftenant.DBHandle(func(tenantId string) (gftenant.Model, error) {
	//		return gftenant.Model{
	//			Debug:     true,
	//			DbName:    tenantId,
	//			DbAddr:    "localhost",
	//			DbPort:    "3306",
	//			DbAccount: "root",
	//			DbPass:    "123456",
	//		}, nil
	//	})
	//}

	r.Middleware.Next()
}
