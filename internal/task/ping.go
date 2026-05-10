package task

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/hibiken/asynq"
	"time"
)

const (
	TypePing = "ping" // task名称
)

type taskPayload struct {
	Ret string `json:"ret"`
}

// NewPingTask 构建任务参数
func NewPingTask(param string) *asynq.Task {
	payload := gjson.MustEncode(taskPayload{Ret: param})
	// asynq.Queue("low")
	// asynq.TaskID("ping_123")
	// 上面上个参数不建议在这里设置，自定义方法里获取不到task的option参数（小写未开放），从而无法删除相同task_id的任务
	return asynq.NewTask(TypePing, payload, asynq.MaxRetry(2))
}

// HandlePingTask 执行任务
func HandlePingTask(ctx context.Context, t *asynq.Task) error {
	var p taskPayload
	if err := gjson.DecodeTo(t.Payload(), &p); err != nil {
		return err
	}
	fmt.Println(p.Ret, time.Now())

	//time.Sleep(5 * time.Second)

	// 这里返回错误会记录到错误日志
	return nil
}
