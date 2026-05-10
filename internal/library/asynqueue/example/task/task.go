package task

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"log"

	"github.com/hibiken/asynq"
)

// A list of task types.
const (
	TypeWelcomeEmail  = "email:welcome"
	TypeReminderEmail = "email:reminder"
)

// Task payload for any email related tasks.
type emailTaskPayload struct {
	// ID for the email recipient.
	UserID int
}

func NewWelcomeEmailTask(id int) (*asynq.Task, error) {
	payload := gjson.MustEncode(emailTaskPayload{UserID: id})
	return asynq.NewTask(TypeWelcomeEmail, payload), nil
}

func NewReminderEmailTask(id int) (*asynq.Task, error) {
	payload := gjson.MustEncode(emailTaskPayload{UserID: id})
	return asynq.NewTask(TypeReminderEmail, payload), nil
}

func HandleWelcomeEmailTask(ctx context.Context, t *asynq.Task) error {
	var p emailTaskPayload
	if err := gjson.DecodeTo(t.Payload(), &p); err != nil {
		return err
	}
	log.Printf(" [*] Send Welcome Email to User %d", p.UserID)
	return gerror.New("could not send email to the user")
}

func HandleReminderEmailTask(ctx context.Context, t *asynq.Task) error {
	var p emailTaskPayload
	if err := gjson.DecodeTo(t.Payload(), &p); err != nil {
		return err
	}
	log.Printf(" [*] Send Reminder Email to User %d", p.UserID)
	return nil
}
