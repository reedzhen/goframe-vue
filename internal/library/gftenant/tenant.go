package gftenant

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

var dbSwitch Handle

func DBHandle(i Handle) {
	dbSwitch = i
}

type Handle func(tenantId string) (Model, error)

var config *Config

type Config struct {
	Mode       string `json:"mode" dc:"none"`            // none 不启用  column 列模式  db 数据库模式
	ColumnName string `json:"columnName" dc:"tenant_id"` // column模式的列名 默认 tenant_id
}

func init() {
	config = &Config{
		Mode: ModeNone,
	}
	v := g.Cfg().MustGet(gctx.GetInitCtx(), "tenant")
	if !v.IsEmpty() {
		if err := v.Scan(&config); err != nil {
			panic(err)
		}
	}
}

func GetColumnName() string {
	if config.ColumnName == "" {
		return "tenant_id"
	}
	return strings.ToLower(config.ColumnName)
}

// GetTenant 获取当前上下文贴牌ID
func GetTenant(ctx context.Context) int64 {
	return gconv.Int64(ctx.Value(GetColumnName()))
}

// SetTenant 设置当前上下文贴牌ID
func SetTenant(ctx context.Context, tenantId int64) {
	g.RequestFromCtx(ctx).SetCtxVar(GetColumnName(), tenantId)
	//return context.WithValue(ctx, GetColumnName(), tenantId)
}

// Init 初始化当前DB链接，分组和dbname同名
func Init(ctx context.Context) gdb.DB {
	if strings.ToLower(config.Mode) != ModeDB {
		return nil
	}

	tenantId := GetTenant(ctx)
	dbName := fmt.Sprintf("tenant_%d", tenantId)

	m, err := dbSwitch(dbName)
	if err != nil {
		panic(err)
	}
	if ll := gdb.GetConfig(dbName); len(ll) == 0 {
		gdb.SetConfigGroup(dbName, gdb.ConfigGroup{
			gdb.ConfigNode{
				Debug: m.Debug,
				Link:  fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.DbAccount, m.DbPass, m.DbAddr, m.DbPort, m.DbName),
				Type:  "mysql",
				Name:  dbName, // 这里补一个Name,否则sql打印会有问题
			},
		})
	}
	err = g.DB(dbName).PingMaster()
	if err != nil {
		panic(err)
	}

	return g.DB(dbName)
}
