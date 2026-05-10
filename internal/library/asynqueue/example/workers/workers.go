package main

import (
	"github.com/hibiken/asynq"
	"goframe-vben/internal/library/asynqueue/example/task"
	"log"
)

// go run workers/workers.go
func main() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "192.168.2.5:6379"},
		asynq.Config{Concurrency: 10},
	)
	mux := asynq.NewServeMux()

	mux.HandleFunc(task.TypeWelcomeEmail, task.HandleWelcomeEmailTask)
	mux.HandleFunc(task.TypeReminderEmail, task.HandleReminderEmailTask)

	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}
