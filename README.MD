# GoFrame-VUE

GoFrame-VUE 是一个基于 GoFrame 2 的后台服务基础框架，面向中后台系统和后续业务项目的二次开发。

项目已经提供基础 RBAC、组织与数据权限、动态菜单、租户、配置中心、日志、定时任务、文件上传、异步任务和插件示例。新项目可以在这些通用能力之上开发业务模块，不需要重复搭建后台基础设施。

配套管理端前端项目：`/Users/reedzhen/Code/vue3/goframe-vibe-admin`

## 技术栈

- GoFrame 2
- MySQL
- Redis
- GToken
- Asynq
- golang-migrate

## 已有能力

### 权限与组织

- 用户管理：用户增删改查、状态管理、密码重置。
- 角色管理：角色维护、菜单授权、数据范围设置。
- 菜单管理：动态菜单、按钮权限标识、接口权限校验。
- 组织管理：树形组织结构与上下级关系。
- 登录认证：验证码、登录限流、Token 校验、全局接口限流。

### 系统能力

- 字典管理：字典类型和字典数据维护。
- 配置中心：按模块维护系统、上传和支付配置。
- 操作日志：记录和查询后台操作。
- 登录日志：查询登录记录。
- 定时任务：任务维护、启停、单次执行和执行日志。
- 文件上传：支持本地存储和阿里云 OSS。
- OpenAPI 文档：启动后访问 `/swagger`。

### 扩展能力

- 租户与租户套餐：租户维护、套餐菜单授权和权限同步。
- 异步任务：基于 Asynq 和 Redis 的任务队列。
- 插件示例：`addons/example` 展示插件路由、服务和异步任务注册方式。
- 代码生成：提供 GoFrame DAO、Controller、Service 和项目内页面模板生成基础。
- 支付基础：包含微信支付、扫呗支付、支付订单和退款基础代码。该模块仍需完成生产化改造后再用于正式业务。

## 目录结构

```text
api/                  接口定义
addons/               插件模块
internal/controller/  控制器
internal/logic/       业务逻辑
internal/library/     通用基础库
internal/router/      路由注册
manifest/config/      本地运行配置
resource/migrations/  数据库迁移脚本
```

## 本地启动

### 1. 准备依赖

- Go 1.23 或更高版本
- MySQL
- Redis

### 2. 准备配置

维护本地文件 `manifest/config/config.yaml`：

- `database.default.link`：MySQL 连接。
- `redis.default`：Redis 地址、库编号和密码。
- `app.mode`：本地使用 `develop`，正式环境使用 `product`。
- `app.mockEnable`：仅允许本地调试开启，正式环境必须关闭。

`manifest/config/config.yaml` 已被 `.gitignore` 忽略。不要提交密码、Token、证书或生产连接信息。

### 3. 初始化数据库

迁移脚本位于 `resource/migrations`。当前首个迁移 `000001_init_db.up.sql` 包含 `DROP TABLE IF EXISTS`，正式库和已有数据库不得直接开启自动迁移。

新建空数据库后，先人工审查 SQL，再手动执行：

```bash
migrate -path resource/migrations \
  -database 'mysql://root:password@tcp(127.0.0.1:3306)/goframe-vue?charset=utf8mb4' \
  up
```

老库接入、版本基线和回滚说明见 [迁移文档](internal/library/migrate/README.MD)。

### 4. 启动服务

```bash
go run main.go
```

默认服务地址为 `http://127.0.0.1:8009`：

- Swagger：`http://127.0.0.1:8009/swagger`
- OpenAPI JSON：`http://127.0.0.1:8009/api.json`

## 常用命令

```bash
# 运行测试
GOCACHE=/private/tmp/goframe-vue-gocache go test ./...

# 生成 DAO、DO 和 Entity
make dao

# 生成 Controller
make ctrl

# 生成 Service
make service

# 构建
make build
```

## 正式发布前检查

- 使用独立的生产配置，确认 `app.mode=product`、`app.mockEnable=false`。
- 更换数据库、Redis、GToken、对象存储和支付相关密钥，不提交到仓库。
- 清理并重新制作初始化 migration：移除历史用户、手机号、测试 Token、开发地址和业务测试数据。
- 人工审查全部 migration；正式库首次接入时先建立基线，不直接执行带 `DROP TABLE IF EXISTS` 的初始化脚本。
- 补齐登录密码校验、登录日志和认证回归测试。
- 审计全部 `noAuth:"true"` 路由，只保留确实需要跳过角色权限校验的接口。
- 明确租户模式；需要多租户时，补齐并启用租户中间件。
- 确认 Redis 可用。当前 HTTP 服务启动时会同时启动 Asynq Worker。
- 完成支付模块生产化改造前，不开放支付和退款能力；支付回调必须使用独立的匿名路由并保留验签和幂等校验。
- 根据部署方式配置反向代理、HTTPS、日志保留、数据库备份和监控告警。
- 执行后端测试和前端生产构建，再做登录、菜单、角色授权、数据权限和文件上传回归。

## 后续完善建议

参考 [HotGo V2](https://github.com/bufanyun/hotgo/tree/v2.0) 的设计，建议按优先级逐步完善：

1. 发布基线：清理初始化数据，补生产配置模板、部署文档、版本发布流程和回归清单。
2. 认证闭环：补齐密码校验、登录成功或失败日志、租户中间件启用策略和安全测试。
3. 插件生命周期：在现有 `addons` 注册机制上补插件创建、安装、更新、卸载和数据库脚本管理。
4. 代码生成：将现有模板扩展为可配置的 CURD、树表和前后端联动生成。
5. 运维能力：增加服务监控、在线用户、服务日志查询和告警。
6. 通知能力：增加通知公告、站内信和 WebSocket 推送。
7. 存储能力：在本地和 OSS 之外，按业务需要增加 MinIO、COS、七牛云等驱动。
8. 支付中心：将支付配置、回调验签、退款状态同步、幂等处理和密钥管理整理为可复用模块。

## 参考项目

- [HotGo V2](https://github.com/bufanyun/hotgo/tree/v2.0)
- [GoFrame](https://github.com/gogf/gf)
