package main

import (
	"github.com/hibiken/asynq"
	"goframe-vben/internal/library/asynqueue/example/task"
	"log"
	"time"
)

// 先运行client.go 在运行workers.go ; go run client/client.go
func main() {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "192.168.2.5:6379"})

	t1, err := task.NewWelcomeEmailTask(42)
	if err != nil {
		log.Fatal(err)
	}

	t2, err := task.NewReminderEmailTask(42)
	if err != nil {
		log.Fatal(err)
	}

	info, err := client.Enqueue(t1)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(" [*] Successfully enqueued task: %+v", info)

	info, err = client.Enqueue(t2, asynq.ProcessIn(24*time.Hour))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(" [*] Successfully enqueued task: %+v", info)
}
