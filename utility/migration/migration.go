package migration

import (
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

//// Migrate 执行迁移
//func Migrate(ctx context.Context, dbGroup gdb.ConfigGroup) (err error) {
//	if len(dbGroup) == 0 {
//		return
//	}
//	for _, db := range dbGroup {
//		if err = MigrateFile(db); err != nil {
//			g.Log().Error(ctx, "["+db.Name+"]", err)
//			return
//		}
//	}
//	return
//}

// MigrateFile 执行sql文件迁移
func MigrateFile(dbCfg gdb.ConfigNode) (err error) {
	link := fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s?%s", dbCfg.User, dbCfg.Pass, dbCfg.Host, dbCfg.Port, dbCfg.Name, dbCfg.Extra)
	m, err := migrate.New("file://migrations", link)
	if err != nil {
		return
	}
	return m.Up()
}

//// MigrateRes 执行资源（gf resource）迁移
//func MigrateRes(ctx context.Context, dbName string) (err error) {
//	d, err := gfres.New("migrations") // gres.Dump() 路径
//	if err != nil {
//		return
//	}
//
//	dbCfg := g.DB("default").GetConfig()
//	link := fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s?%s", dbCfg.User, dbCfg.Pass, dbCfg.Host, dbCfg.Port, dbName, dbCfg.Extra)
//	m, err := migrate.NewWithSourceInstance("gfres", d, link)
//	if err != nil {
//		return
//	}
//
//	// 执行迁移（第一次会生成 schema_migrations 表，用来记录迁移版本）
//	if err = m.Up(); err != nil {
//		if err == migrate.ErrNoChange {
//			return errors.New("[" + dbName + "] no change")
//		}
//	}
//	return
//}

//func Migration(dbName string) (err error) {
//	d, err := gfres.New("migrations") // gres.Dump() 路径
//	if err != nil {
//		return
//	}
//
//	dbCfg := g.DB().GetConfig()
//	link := fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s?%s", dbCfg.User, dbCfg.Pass, dbCfg.Host, dbCfg.Port, dbName, dbCfg.Extra)
//	m, err := migrate.NewWithSourceInstance("gfres", d, link)
//	if err != nil {
//		return
//	}
//
//	if err = m.Up(); err != nil {
//		return
//	}
//	g.Dump("migrate success")
//
//	return
//}
