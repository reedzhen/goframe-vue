package plugin

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"goframe-vben/internal/service"
)

type Cron struct {
}

func (p *Cron) Name() string {
	return "cron"
}

func (p *Cron) Author() string {
	return "Larry Liu"
}

func (p *Cron) Version() string {
	return "1.0"
}

func (p *Cron) Description() string {
	return "cron service"
}

func (p *Cron) Install(s *ghttp.Server) error {
	var ctx = gctx.New()

	// @every 4h 每4个小时执行一次
	// @every 1m 每分钟执行一次
	// 秒 分 时  日 月 周
	// 0 05 17 * * ?
	//if gmode.Mode() != gmode.DEVELOP {
	service.Cron().Start(ctx)
	//}

	return nil
}

func (p *Cron) Remove() error {
	service.Cron().Remove()
	return nil
}
