package clnhttp

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

type ClnHttp struct {
	appId    string
	secret   string
	tenantId string
}

func NewClnHttp(appId string, secret string, tenantId string) *ClnHttp {
	return &ClnHttp{appId: appId, secret: secret, tenantId: tenantId}
}

// ClnResponse 接口返回
type ClnResponse struct {
	Code    int         `json:"code"`    // 错误码((0:成功, 1:失败, >1:错误码))
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据(业务接口定义具体数据结构)
}

func (m *ClnHttp) Get(ctx context.Context, url string, data ...interface{}) (res *ClnResponse, err error) {
	return m.HttpRequest(ctx, "GET", url, data...)
}

func (m *ClnHttp) Post(ctx context.Context, url string, data ...interface{}) (res *ClnResponse, err error) {
	return m.HttpRequest(ctx, "POST", url, data...)
}

func (m *ClnHttp) HttpRequest(ctx context.Context, method string, url string, data ...interface{}) (res *ClnResponse, err error) {
	headerMap := make(map[string]string)
	headerMap["AppId"] = m.appId
	headerMap["X-STORE"] = m.tenantId

	err = g.Client().ContentJson().Header(headerMap).RequestVar(ctx, method, url, data...).Scan(&res)
	return
}
