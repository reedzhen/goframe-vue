package ipcity

import (
	"context"
	"fmt"
	"goframe-vben/internal/library/cache"
	"net"
	"net/url"
	"time"

	"github.com/gogf/gf/v2/encoding/gcharset"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

const (
	WhoisApi       = "https://whois.pconline.com.cn/ipJson.jsp?json=true&ip=%s"
	LocalhostIPv4  = "127.0.0.1" // LocalhostIPv4 和 LocalhostIPv6 是本地回环IP地址
	LocalhostIPv6  = "::1"
	CacheIPCity    = "cache:ip_city:%s" // IP归属地缓存
	CacheIPCityTTL = 24 * 60 * 60       // IP归属地缓存缓存时间（s）
)

type WhoisData struct {
	Ip         string `json:"ip"`
	Pro        string `json:"pro" `
	ProCode    string `json:"proCode" `
	City       string `json:"city" `
	CityCode   string `json:"cityCode"`
	Region     string `json:"region"`
	RegionCode string `json:"regionCode"`
	Addr       string `json:"addr"`
	Err        string `json:"err"`
}

// GetWhoisLocation 通过Whois查询IP归属地
func GetWhoisLocation(ctx context.Context, ip string) (out *WhoisData, err error) {

	queryUrl := fmt.Sprintf(WhoisApi, url.QueryEscape(ip))
	response, err := g.Client().Timeout(5*time.Second).Retry(3, time.Second).Get(ctx, queryUrl)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Close()
	}()

	// 转换字符集
	str, err := gcharset.ToUTF8("GBK", response.ReadAllString())
	if err != nil {
		return nil, err
	}

	if err = gconv.Scan([]byte(str), &out); err != nil {
		return nil, err
	}
	return
}

// GetLocation 获取IP归属地信息
func GetLocation(ctx context.Context, ip string) (data *WhoisData, err error) {
	if net.ParseIP(ip) == nil {
		return nil, gerror.Newf("无效IP:%v", ip)
	}

	// 检查是否为本地回环地址
	if ip == LocalhostIPv4 || ip == LocalhostIPv6 {
		return &WhoisData{Ip: LocalhostIPv4, City: "内网IP"}, nil
	}

	// 获取缓存
	cacheKey := fmt.Sprintf(CacheIPCity, ip)
	cacheData, err := cache.Instance().Get(ctx, cacheKey)
	if err != nil {
		return
	}
	if !cacheData.IsNil() {
		_ = cacheData.Scan(&data)
		return
	}

	// 通过Whois查询IP归属地
	data, err = GetWhoisLocation(ctx, ip)
	if err != nil || data == nil {
		return
	}

	// 设置缓存
	if err = cache.Instance().Set(ctx, cacheKey, data, CacheIPCityTTL*time.Second); err != nil {
		return
	}

	return
}
