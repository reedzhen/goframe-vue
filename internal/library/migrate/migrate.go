package migrate

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	gomigrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	defaultSource        = "file"
	defaultPath          = "resource/migrations"
	defaultDatabaseGroup = "default"

	sourceFile  = "file"
	sourceGfRes = "gfres"
)

// Config 数据库迁移配置。
type Config struct {
	Enabled       bool   `json:"enabled"`       // 是否在项目启动时自动执行迁移
	Source        string `json:"source"`        // 迁移文件来源：file 本地目录，gfres 打包资源
	Path          string `json:"path"`          // 迁移文件目录
	DatabaseGroup string `json:"databaseGroup"` // 数据库配置组
}

// DefaultConfig 返回默认迁移配置。
func DefaultConfig() Config {
	return Config{
		Enabled:       false,
		Source:        defaultSource,
		Path:          defaultPath,
		DatabaseGroup: defaultDatabaseGroup,
	}
}

// Run 按配置执行数据库迁移，适合在项目启动阶段调用。
func Run(ctx context.Context) error {
	cfg, err := LoadConfig(ctx)
	if err != nil {
		return err
	}
	if !cfg.Enabled {
		g.Log().Info(ctx, "数据库迁移已关闭")
		return nil
	}

	if err = Up(ctx, cfg); err != nil {
		return err
	}
	return nil
}

// LoadConfig 读取迁移配置，并补齐默认值。
func LoadConfig(ctx context.Context) (Config, error) {
	cfg := DefaultConfig()
	v := g.Cfg().MustGet(ctx, "migration")
	if v == nil || v.IsNil() {
		return cfg, nil
	}
	if err := v.Scan(&cfg); err != nil {
		return cfg, err
	}
	normalizeConfig(&cfg)
	return cfg, nil
}

// Up 执行所有未应用的 up 迁移。
func Up(ctx context.Context, cfg Config) error {
	m, err := New(ctx, cfg)
	if err != nil {
		return err
	}
	defer closeMigrate(ctx, m)

	if err = m.Up(); err != nil {
		if errors.Is(err, gomigrate.ErrNoChange) {
			g.Log().Info(ctx, "数据库迁移无变更")
			return nil
		}
		return err
	}

	g.Log().Info(ctx, "数据库迁移执行完成")
	return nil
}

// New 创建 migrate 实例，调用方可用于 Version、Steps、Force 等高级操作。
func New(ctx context.Context, cfg Config) (*gomigrate.Migrate, error) {
	normalizeConfig(&cfg)

	dbURL, err := buildDatabaseURL(ctx, cfg.DatabaseGroup)
	if err != nil {
		return nil, err
	}

	switch cfg.Source {
	case sourceFile:
		return gomigrate.New("file://"+cfg.Path, dbURL)
	case sourceGfRes:
		driver, err := newGfResourceSource(cfg.Path)
		if err != nil {
			return nil, err
		}
		return gomigrate.NewWithSourceInstance(sourceGfRes, driver, dbURL)
	default:
		return nil, gerror.Newf("不支持的迁移文件来源: %s", cfg.Source)
	}
}

func normalizeConfig(cfg *Config) {
	if cfg.Source == "" {
		cfg.Source = defaultSource
	}
	cfg.Source = strings.ToLower(cfg.Source)
	if cfg.Path == "" {
		cfg.Path = defaultPath
	}
	if cfg.DatabaseGroup == "" {
		cfg.DatabaseGroup = defaultDatabaseGroup
	}
}

func buildDatabaseURL(ctx context.Context, group string) (string, error) {
	db := g.DB(group)
	if err := db.PingMaster(); err != nil {
		return "", err
	}

	node := db.GetConfig()
	if node == nil {
		return "", gerror.Newf("数据库配置组不存在: %s", group)
	}
	if node.Name == "" {
		return "", gerror.New("数据库名不能为空")
	}

	switch strings.ToLower(node.Type) {
	case "mysql", "mariadb":
		return mysqlURL(*node), nil
	case "pgsql", "postgres", "postgresql":
		return postgresURL(*node), nil
	default:
		return "", gerror.Newf("数据库迁移暂不支持当前类型: %s", node.Type)
	}
}

func mysqlURL(node gdb.ConfigNode) string {
	if node.Protocol == "" {
		node.Protocol = "tcp"
	}
	if node.Port == "" {
		node.Port = "3306"
	}

	auth := ""
	if node.User != "" {
		auth = url.QueryEscape(node.User)
		if node.Pass != "" {
			auth += ":" + url.QueryEscape(node.Pass)
		}
		auth += "@"
	}

	extra := node.Extra
	if extra == "" {
		extra = "charset=utf8mb4&parseTime=True&loc=Local"
	}
	if !strings.Contains(extra, "multiStatements=") {
		extra += "&multiStatements=true"
	}

	return fmt.Sprintf(
		"mysql://%s%s(%s:%s)/%s?%s",
		auth,
		node.Protocol,
		node.Host,
		node.Port,
		node.Name,
		extra,
	)
}

func postgresURL(node gdb.ConfigNode) string {
	if node.Host == "" {
		node.Host = "localhost"
	}
	if node.Port == "" {
		node.Port = "5432"
	}

	u := url.URL{
		Scheme: "postgres",
		Host:   node.Host + ":" + node.Port,
		Path:   "/" + node.Name,
	}
	if node.User != "" {
		if node.Pass != "" {
			u.User = url.UserPassword(node.User, node.Pass)
		} else {
			u.User = url.User(node.User)
		}
	}

	query, _ := url.ParseQuery(node.Extra)
	if query.Get("sslmode") == "" {
		query.Set("sslmode", "disable")
	}
	if query.Get("x-multi-statement") == "" {
		query.Set("x-multi-statement", "true")
	}
	u.RawQuery = query.Encode()

	return u.String()
}

func closeMigrate(ctx context.Context, m *gomigrate.Migrate) {
	if sourceErr, databaseErr := m.Close(); sourceErr != nil || databaseErr != nil {
		g.Log().Warningf(ctx, "关闭数据库迁移资源失败 source=%v database=%v", sourceErr, databaseErr)
	}
}
