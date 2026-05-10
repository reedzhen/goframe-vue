### migrate 新建迁移脚本文件

```shell
# ext 后缀, dir 生成的文件位置, seq 使用数字前缀（默认时间戳）
migrate create -ext sql -dir resource/migrations -seq init_db
```

### 项目启动时自动迁移

迁移 SQL 统一放在 `resource/migrations`，不要放到 `manifest`。`manifest` 主要放配置和部署清单，迁移脚本属于运行资源，放在 `resource` 更清晰。

```yaml
migration:
  enabled: false # 确认迁移脚本安全后再打开
  source: "file" # file读取本地目录；gfres读取gf pack打包资源
  path: "resource/migrations"
  databaseGroup: "default"
```

迁移封装会根据 GoFrame 数据库配置的 `type` 自动选择驱动，目前支持 `mysql/mariadb` 和 `pgsql/postgres/postgresql`。

当前 `000001_init_db.up.sql` 包含 `DROP TABLE IF EXISTS`，老库第一次接入时建议先用 `force` 建立基线版本，再打开启动迁移，避免误清空已有表。

### migrate 手动执行迁移文件

```shell
# 1.不能有“&”这个符号 2.注意数据库连接写法 mysql://*
migrate -path resource/migrations -database mysql://root:xxx@tcp(db_url:3306)/db_name?charset=utf8mb4 up 1
```

### migrate 回滚迁移文件

```shell
# 1.不能有“&”这个符号 2.注意数据库连接写法 mysql://*
migrate -path resource/migrations -database mysql://root:xxx@tcp(db_url:3306)/db_name?charset=utf8mb4 down 1
```

### migrate 删除数据库中的所有内容 用于在重新初始化数据库或进行大规模迁移之前清空数据库中的所有表和其他对象

```shell
migrate -path resource/migrations -database mysql://root:xxx@tcp(db_url:3306)/db_name?charset=utf8mb4 drop
```

### migrate 打印迁移版本号

```shell
migrate -path resource/migrations -database mysql://root:xxx@tcp(db_url:3306)/db_name?charset=utf8mb4 version
```

### migrate 设置迁移版本号，不执行迁移，只设置schema_migrations表迁移版本号(ignores dirty state)，报错修复版本时用

例如：当前执行up报错了，schema_migrations表，version+N 和 dirty+1
1. 此时我们需要修复 sql
2. 执行 force V 命令，回滚到之前的版本（我之前不知道这个指令，都是在数据库中手动修复的）
3. 重新执行 up 指令

```shell
migrate -path resource/migrations -database mysql://root:xxx@tcp(db_url:3306)/db_name?charset=utf8mb4 force 版本号
```

### migrate 迁移至某个版本(up是基于当前表schema_migrations最新版本号 +N 升级,goto V 这个V必须存在)

```shell
migrate -path resource/migrations -database mysql://root:xxx@tcp(db_url:3306)/db_name?charset=utf8mb4 goto 版本号
```

### migrate 其他用法，可以打包迁移文件到 packed.go，然后在项目启动的时候执行迁移（缺点是打包的时候容易忘）

```shell
gf pack resource/migrations internal/packed/packed.go -y
```