package middleware

import (
	"mime"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gmode"
	"goframe-vben/internal/codes"
	"goframe-vben/utility/response"
)

const (
	contentTypeEventStream  = "text/event-stream"
	contentTypeOctetStream  = "application/octet-stream"
	contentTypeMixedReplace = "multipart/x-mixed-replace"
)

var (
	// streamContentType 用于判断是否是流式响应
	streamContentType = []string{
		contentTypeEventStream,
		contentTypeOctetStream,
		contentTypeMixedReplace,
	}
)

// HandlerResponse 自定义统一响应返回中间件
func (s *sMiddleware) HandlerResponse(r *ghttp.Request) {
	r.Middleware.Next()

	// 如果有自定义的返回内容，或者已经写入了数据，则当前中间件直接退出
	if r.Response.BufferLength() > 0 || r.Response.BytesWritten() > 0 {
		return
	}

	// 如果是流式响应，不走统一响应格式
	mediaType, _, _ := mime.ParseMediaType(r.Response.Header().Get("Content-Type"))
	for _, ct := range streamContentType {
		if mediaType == ct {
			return
		}
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
		// 处理 HTTP Status 非 200 的情况，例如 404 等
		if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
			switch r.Response.Status {
			case http.StatusNotFound:
				code = codes.CodeNotFound
			case http.StatusForbidden:
				code = codes.CodeForbidden
			default:
				code = codes.CodeUnknown
			}
			msg := code.Message()
			err = gerror.NewCode(code, msg)
			r.SetError(err)

			response.JsonExit(r, code.Code(), msg)
		} else {
			response.JsonExit(r, code.Code(), code.Message(), res)
		}
	}
}
