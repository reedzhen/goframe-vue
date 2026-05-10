package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gmode"
	"goframe-vben/internal/codes"
	"goframe-vben/utility/response"
)

// ResponseHandler 返回处理中间件
func (s *sMiddleware) ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()

	// 如果已经有返回内容，那么该中间件什么也不做
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = codes.CodeOK
	)
	if err != nil {
		code = gerror.Code(err)
		if code == gcode.CodeNil {
			code = gcode.CodeValidationFailed
		}

		msg := ""
		switch code {
		case gcode.CodeDbOperationError, gcode.CodeInternalError:
			// db错误不可以把sql暴露给前端
			msg = "服务器居然开小差了，请稍后再试吧！"
			code = codes.CodeInternalServerError
		case gcode.CodeValidationFailed:
			// 请求校验的错误默认都会被记录到server错误日志中，r.SetError(nil)会导致记录操作日志后置中间件获取不到错误信息
			// 这里折中方案是写入错误日志，但是不记录错误堆栈，这种直接抛给用户看的错误记录堆栈没有意义

			// 方案1
			// r.SetError(nil)

			// 方案2
			if !gmode.IsDevelop() {
				r.SetError(gerror.NewWithOption(gerror.Option{
					Stack: false,
					Text:  err.Error(),
					Code:  gcode.CodeValidationFailed,
				}))
			}
			code = codes.CodeBadRequest
			msg = err.Error()
		case gcode.CodeNotImplemented:
			code = codes.CodeNotImplemented
			msg = codes.CodeNotImplemented.Message()
		default:
			msg = err.Error()
		}
		response.JsonExit(r, code.Code(), msg)
	} else {
		response.JsonExit(r, code.Code(), code.Message(), res)
	}
}
