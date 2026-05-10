package user

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"goframe-vben/addons/example/api/user/admin"
	"goframe-vben/addons/example/task"
	"goframe-vben/internal/library/asynqueue"
)

func (c *ControllerAdmin) Task(ctx context.Context, req *admin.TaskReq) (res *admin.TaskRes, err error) {
	taskInfo, err := asynqueue.Client.EnqueueContext(ctx, task.NewHelloTask("example_task_test"))
	if err != nil {
		return nil, err
	}
	g.Dump(taskInfo)
	return
}
