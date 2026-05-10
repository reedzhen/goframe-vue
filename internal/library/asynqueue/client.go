package asynqueue

import (
	"context"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/hibiken/asynq"
	"time"
)

// Client 队列客户端
var Client *client

type client struct {
	*asynq.Client
	*asynq.Inspector
}

func NewClient(config gredis.Config) {
	opt := asynq.RedisClientOpt{
		Addr:     config.Address,
		Password: config.Pass,
		DB:       config.Db,
	}
	Client = &client{Client: asynq.NewClient(opt), Inspector: asynq.NewInspector(opt)}
}

// EnqueueContext 将任务放入队列（任务属于某一个队列)
// asynq.Queue("default") 队列名称，默认default
// asynq.TaskID("order_123") 任务ID，默认随机生成
// asynq.ProcessIn(5*time.Second) 延迟5s后执行
// asynq.Retention(7*24*time.Hour) 完成的任务在指定的时长过后将被删除
// asynq.MaxRetry(maxRetry) 默认错误重试次数 25
// asynq.Unique(time.Duration) 在参数时间返回内是唯一的
func (c *client) EnqueueContext(ctx context.Context, task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	//// 删除重名任务
	//if err := c.delCompletedTask(opts); err != nil {
	//	return nil, err
	//}

	// unique 设置方式
	// Task Type :asynq.Unique(time.Duration)
	// Task Payload : 参数
	// Queue Name : asynq.Queue("default")

	// 默认完成的任务7天后删除
	if !c.hasOption(asynq.RetentionOpt, opts) {
		opts = append(opts, asynq.Retention(7*24*time.Hour))
	}

	return c.Client.EnqueueContext(ctx, task, opts...)
}

func (c *client) hasOption(option asynq.OptionType, opts []asynq.Option) bool {
	has := false
	for _, v := range opts {
		if v.Type() == option {
			has = true
		}
	}
	return has
}

//// delCompletedTask 删除已完成task，不需要了 可以通过unique来实现任务唯一
//func (c *client) delCompletedTask(opts []asynq.Option) error {
//	if len(opts) == 0 {
//		return nil
//	}
//
//	queueName := "default"
//	taskId := ""
//	for _, v := range opts {
//		if v.Type() == asynq.QueueOpt {
//			queueName = gconv.String(v.Value())
//		}
//		if v.Type() == asynq.TaskIDOpt {
//			taskId = gconv.String(v.Value())
//		}
//	}
//
//	// ==空，说明没有手动命名任务名称，会自动使用随机字符串名称，不存在两个相同的任务
//	if taskId == "" {
//		return nil
//	}
//
//	// 删除 task
//	if err := c.Inspector.DeleteTask(queueName, taskId); err != nil {
//		return err
//	}
//	return nil
//}
