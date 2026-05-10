package middleware

import (
	"goframe-vben/internal/codes"
	"goframe-vben/internal/library/contexts"
	"goframe-vben/internal/library/gtoken/backend"
	"goframe-vben/internal/model"
	"goframe-vben/internal/model/entity"
	"goframe-vben/utility/response"

	"github.com/goflyfox/gtoken/v2/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

// CheckJwtToken token校验中间件
func (s *sMiddleware) CheckJwtToken(r *ghttp.Request) {
	gToken := backend.InitGToken()

	// 注册gToken中间件
	middlewareAuth := gtoken.NewDefaultMiddleware(gToken)
	// token校验失败后的返回方法
	middlewareAuth.ResFun = func(r *ghttp.Request, err error) {
		r.Response.WriteJson(g.Map{
			"code":    codes.CodeUnauthorized.Code(),
			"message": "身份认证过期，请重新登录:" + err.Error(),
			"data":    []interface{}{},
		})
		return
	}
	//middlewareAuth.Auth(r)

	// 以下代码重写上面gtokenk框架 middlewareAuth.Auth(r) 方法 begin ---
	if middlewareAuth.HasExcludePath(r) {
		// 如果不需要认证，继续
		r.Middleware.Next()
		return
	}

	// 获取请求token
	token, err := gtoken.GetRequestToken(r)
	if err != nil {
		middlewareAuth.ResFun(r, err)
		return
	}

	userKey, err := middlewareAuth.Token.Validate(r.Context(), token)
	if err != nil {
		middlewareAuth.ResFun(r, err)
		return
	}
	r.SetCtxVar(gtoken.KeyUserKey, userKey)
	// end ---

	_, data, err := gToken.ParseToken(r.Context(), token)
	if err != nil {
		//r.Response.WriteJson(RespError(err))
		response.JsonExit(r, codes.CodeUnauthorized.Code(), err.Error())
	}

	var user entity.SysUser
	if err := gconv.Struct(data, &user); err != nil {
		response.JsonExit(r, codes.CodeUnauthorized.Code(), "授权过期")
	}

	// 设置上下文用户
	contexts.SetUser(r.GetCtx(), &model.Identity{
		UserId:         user.Id,
		Username:       user.Username,
		Nickname:       user.Nickname,
		RoleId:         user.RoleId,
		IsAdmin:        user.IsAdmin == 1,
		OrganizationId: user.OrganizationId,
	})

	r.Middleware.Next()
}
