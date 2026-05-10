package eventbus

import (
	"context"
	"testing"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
)

type testEvent struct {
	ID int
}

func TestBusSubscribeAndPublish(t *testing.T) {
	ctx := context.Background()
	bus := New()
	called := 0

	err := bus.Subscribe("user.created", func(ctx context.Context, event any) error {
		e := event.(*testEvent)
		if e.ID != 1001 {
			t.Fatalf("ID = %d, want 1001", e.ID)
		}
		called++
		return nil
	})
	if err != nil {
		t.Fatalf("Subscribe failed: %v", err)
	}

	if err = bus.Publish(ctx, "user.created", &testEvent{ID: 1001}); err != nil {
		t.Fatalf("Publish failed: %v", err)
	}
	if called != 1 {
		t.Fatalf("called = %d, want 1", called)
	}
}

func TestBusPublishCallsAllHandlersAndReturnsFirstError(t *testing.T) {
	ctx := context.Background()
	bus := New()
	wantErr := gerror.New("failed")
	called := 0

	_ = bus.Subscribe("order.paid", func(ctx context.Context, event any) error {
		called++
		return nil
	})
	_ = bus.Subscribe("order.paid", func(ctx context.Context, event any) error {
		called++
		return wantErr
	})
	_ = bus.Subscribe("order.paid", func(ctx context.Context, event any) error {
		called++
		return nil
	})

	if err := bus.Publish(ctx, "order.paid", "payload"); err != wantErr {
		t.Fatalf("error = %v, want %v", err, wantErr)
	}
	if called != 3 {
		t.Fatalf("called = %d, want 3", called)
	}
}

func TestBusPublishAsync(t *testing.T) {
	ctx := context.Background()
	bus := New()
	done := make(chan struct{})

	_ = bus.Subscribe("async.task", func(ctx context.Context, event any) error {
		close(done)
		return nil
	})

	bus.PublishAsync(ctx, "async.task", "test data")

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("async handler was not called")
	}
}

func TestBusClear(t *testing.T) {
	ctx := context.Background()
	bus := New()
	called := 0

	_ = bus.Subscribe("test.topic", func(ctx context.Context, event any) error {
		called++
		return nil
	})
	bus.Clear()

	_ = bus.Publish(ctx, "test.topic", "data")

	if called != 0 {
		t.Fatalf("called = %d, want 0", called)
	}
}
