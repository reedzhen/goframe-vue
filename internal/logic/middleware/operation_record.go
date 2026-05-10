package middleware

import (
	"encoding/json"
	"fmt"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"

	"goframe-vben/internal/consts"
	"goframe-vben/internal/library/contexts"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"
	"goframe-vben/utility/response"
)

const maxOperationRecordResultLength = 60000

// OperationRecord 记录操作日志
func (s *sMiddleware) OperationRecord(r *ghttp.Request) {
	r.Middleware.Next()

	if r.Method != "POST" {
		return
	}

	in, ok := s.operationRecordInput(r)
	if !ok {
		return
	}

	service.OperationRecord().AsyncCreate(r.GetCtx(), in)
}

func (s *sMiddleware) operationRecordInput(r *ghttp.Request) (dto.OperationRecordCreateInput, bool) {
	ctx := r.Context()
	user := contexts.GetUser(ctx)
	if user == nil {
		return dto.OperationRecordCreateInput{}, false
	}

	var (
		module  string
		summary string
	)
	if ctxData := contexts.GetData(ctx); ctxData != nil {
		module = gstr.Trim(gconv.String(ctxData["tags"]))
		summary = gstr.Trim(gconv.String(ctxData["summary"]))
	}

	var (
		param      = r.GetBodyString()
		remark     string
		jsonResult = r.Response.BufferString()
		errorMsg   string
		status     = consts.OperationRecordStatusSuc
	)
	if gstr.Contains(r.GetHeader("Content-Type"), "multipart/form-data") {
		param = ""
		remark = fmt.Sprintf("上传文件大小: %.2fMB", bytesToMB(len(r.GetBodyString())))
	}
	if len(jsonResult) > maxOperationRecordResultLength {
		jsonResult = ""
	}
	if err := r.GetError(); err != nil {
		errorMsg = err.Error()
		status = consts.OperationRecordStatusErr
	} else {
		var res response.JsonRes
		if err := json.Unmarshal([]byte(jsonResult), &res); err == nil && res.Code != 0 {
			errorMsg = res.Message
			status = consts.OperationRecordStatusErr
		}
	}

	return dto.OperationRecordCreateInput{
		UserId:     user.UserId,
		Username:   user.Username,
		Nickname:   user.Nickname,
		Url:        r.URL.Path,
		Method:     r.Method,
		Module:     module,
		Summary:    summary,
		Param:      param,
		JsonResult: jsonResult,
		ErrorMsg:   errorMsg,
		SpendTime:  gtime.Now().Sub(r.EnterTime).Milliseconds(),
		TraceId:    gctx.CtxId(ctx),
		Status:     status,
		Platform:   contexts.GetModule(ctx),
		UserAgent:  r.UserAgent(),
		Ip:         r.GetClientIp(),
		Remark:     remark,
	}, true
}

func bytesToMB(bytes int) float64 {
	return float64(bytes) / (1024.0 * 1024.0)
}
