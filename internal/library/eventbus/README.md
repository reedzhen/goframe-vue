# EventBus 使用指南

这是一个进程内事件总线，用来做模块之间的轻量解耦。现在只保留最常用的能力：

- `eventbus.Subscribe(topic, handler)`：订阅事件
- `eventbus.Publish(ctx, topic, event)`：同步发布，等待所有 handler 执行完
- `eventbus.PublishAsync(ctx, topic, event)`：异步发布，立即返回
- `eventbus.New()`：创建独立总线，主要用于测试
- `eventbus.Clear()`：清空默认总线，主要用于测试

## 快速使用

定义事件名和事件数据：

```go
package consts

const EventUserCreated = "user.created"
```

```go
package event

type UserCreated struct {
	UserID   int64
	Username string
}
```

在启动时注册订阅者：

```go
package user

import (
	"context"

	"goframe-vben/internal/consts"
	"goframe-vben/internal/library/eventbus"
	modelevt "goframe-vben/internal/model/event"
)

func init() {
	_ = eventbus.Subscribe(consts.EventUserCreated, onUserCreated)
}

func onUserCreated(ctx context.Context, data any) error {
	e := data.(*modelevt.UserCreated)
	// 处理用户创建后的逻辑，例如发通知、写审计日志、同步插件数据
	_ = e
	return nil
}
```

在业务成功后发布事件：

```go
err := eventbus.Publish(ctx, consts.EventUserCreated, &modelevt.UserCreated{
	UserID:   user.Id,
	Username: user.Username,
})
if err != nil {
	return err
}
```

耗时或不影响主流程的逻辑可以异步发布：

```go
eventbus.PublishAsync(ctx, consts.EventUserCreated, &modelevt.UserCreated{
	UserID:   user.Id,
	Username: user.Username,
})
```

## 使用建议

事件名使用 `资源.动作`，例如 `user.created`、`order.paid`。

事件数据只放必要字段，不要直接传完整实体，避免订阅者依赖太多发布者内部细节。

事务内不要发布事件。先提交事务，确认业务数据落库成功，再 `Publish` 或 `PublishAsync`，否则后续回滚时可能已经触发了订阅者。

同步 `Publish` 会执行所有 handler，并返回第一个错误；即使某个 handler 失败，后面的 handler 仍会继续执行。异步 `PublishAsync` 会忽略 handler 错误，适合发消息、写日志、外部通知这类旁路动作。
