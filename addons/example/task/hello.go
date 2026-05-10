package task

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/hibiken/asynq"
	"time"
)

const (
	TypeHello = "example:hello" // task名称
)

type taskPayload struct {
	Ret string
}

// NewHelloTask 构建任务参数
func NewHelloTask(param string) *asynq.Task {
	payload := gjson.MustEncode(taskPayload{Ret: param})
	return asynq.NewTask(TypeHello, payload, asynq.Queue("low"))
}

// HandleHelloTask 执行任务
func HandleHelloTask(ctx context.Context, t *asynq.Task) error {
	var p taskPayload
	if err := gjson.DecodeTo(t.Payload(), &p); err != nil {
		return err
	}
	time.Sleep(5 * time.Second)
	fmt.Println(p.Ret)
	return nil
}
