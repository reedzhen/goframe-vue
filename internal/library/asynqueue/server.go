package asynqueue

import (
	"context"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hibiken/asynq"
	"goframe-vben/internal/consts"
	"time"
)

var Server *server

type server struct {
	*asynq.Server
	serveMux *asynq.ServeMux
}

type Handle func(ctx context.Context, task *asynq.Task) error

var DefaultConfig = asynq.Config{
	// Concurrency 表示最大并发处理任务数。
	Concurrency: 50,
	// 队列优先级，所有队列都必须在这里申明优先级，否则 task 状态永远是 pending
	Queues: map[string]int{
		"critical": 6, // 严重
		"default":  3, // 默认
		"low":      1, // 低级
	},
	StrictPriority: true,
	// 重试间隔
	RetryDelayFunc: func(n int, e error, t *asynq.Task) time.Duration {
		return time.Duration(fib(n+1)) * time.Second
		//return time.Duration(3) * time.Second
	},
	ErrorHandler: &ReportError{},
}

// NewServer 创建一个新的 asynq 服务器实例
func NewServer(redisCgf gredis.Config) {
	cli := asynq.NewServer(asynq.RedisClientOpt{
		Addr:     redisCgf.Address,
		Password: redisCgf.Pass,
		DB:       redisCgf.Db,
	}, DefaultConfig)

	Server = &server{Server: cli, serveMux: asynq.NewServeMux()}
}

// RegisterHandle 注册处理特定任务函数。这里通过一个中间函数捕获处理过程中的错误并记录日志。
func (s *server) RegisterHandle(topic string, fn Handle) {
	s.serveMux.HandleFunc(topic, fn)
}

// Use 中间件
func (s *server) Use(mws ...asynq.MiddlewareFunc) {
	s.serveMux.Use(mws...)
}

// Start 启动服务
func (s *server) Start() error {
	return s.Server.Start(s.serveMux)
}

// Fibonacci returns successive Fibonacci numbers starting from 1
func fib(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}

//
//func BaseHandle(cb func(ctx context.Context, task *asynq.Task) error) func(ctx context.Context, task *asynq.Task) error {
//	return func(ctx context.Context, task *asynq.Task) error {
//		b := task.Payload()
//		var base Base
//		_ = json.Unmarshal(b, &base)
//
//		//ctx = company.SetCompany(ctx, base.CompanyId)
//		return cb(ctx, task)
//	}
//}

type ReportError struct{}

func (r *ReportError) HandleError(ctx context.Context, task *asynq.Task, err error) {
	retried, _ := asynq.GetRetryCount(ctx)
	maxRetry, _ := asynq.GetMaxRetry(ctx)
	if retried >= maxRetry {
		g.Log(consts.LoggerGroupAsynq).Errorf(ctx, "[task] %s [payload] %v [retried] %d [maxRetry] %d [error] %+v", task.Type(), string(task.Payload()), retried, maxRetry, err)
	} else {
		g.Log(consts.LoggerGroupAsynq).Errorf(ctx, "[task] %s [payload] %v [error] %+v", task.Type(), string(task.Payload()), err)
	}
	// 可以发送邮件啥的
}
