package system

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/mssola/user_agent"
	"goframe-vben/internal/dao"
	"goframe-vben/internal/library/ipcity"
	"goframe-vben/internal/library/query"
	"goframe-vben/internal/model/do"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/model/entity"
	"goframe-vben/internal/service"
	"strings"
)

type sLoginRecord struct{}

func init() {
	service.RegisterLoginRecord(NewLoginRecord())
}

func NewLoginRecord() *sLoginRecord {
	return &sLoginRecord{}
}

// AsyncCreate 异步调用登录日志
func (s *sLoginRecord) AsyncCreate(ctx context.Context, in dto.LoginRecordCreateInput) {
	if err := grpool.AddWithRecover(ctx, func(ctx context.Context) {
		if err := s.Create(ctx, in); err != nil {
			g.Log().Error(ctx, err)
		}
	}, func(ctx context.Context, err error) {
		g.Log().Error(ctx, err)
	}); err != nil {
		g.Log().Error(ctx, err)
	}
	return
}

// Create 新建登录日志
func (s *sLoginRecord) Create(ctx context.Context, in dto.LoginRecordCreateInput) (err error) {
	// 解析 User-Agent
	ua := user_agent.New(g.RequestFromCtx(ctx).UserAgent())
	bName, bVersion := ua.Browser()

	ip := g.RequestFromCtx(ctx).GetClientIp()

	data := do.SysLoginRecord{
		Username:  in.Username,
		Os:        ua.OS(),
		Device:    ua.Platform(),
		Browser:   bName + " " + bVersion,
		Ip:        ip,
		IpCity:    nil,
		LoginType: in.LoginType,
		Remark:    in.Remark,
		//TenantId:  in.TenantId,
	}

	// 获取ip归属地
	location, err := ipcity.GetLocation(ctx, ip)
	if err != nil {
		// IP 归属地是辅助信息，查询失败不能影响登录日志落库。
		g.Log().Warning(ctx, "获取登录IP归属地失败", err)
	}
	if location != nil {
		data.IpCity = strings.TrimSpace(location.Pro + " " + location.City)
	}

	_, err = dao.SysLoginRecord.Ctx(ctx).Data(data).Insert()
	return
}

// Page 登录日志分页
func (s *sLoginRecord) Page(ctx context.Context, in dto.LoginRecordPageInput) (out *query.Result, err error) {
	list := make([]*entity.SysLoginRecord, 0)
	out, err = query.Page(dao.SysLoginRecord.Ctx(ctx), &in, &list)
	return
}
