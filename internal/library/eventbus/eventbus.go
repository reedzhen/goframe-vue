package eventbus

import (
	"context"
	"sync"

	"github.com/gogf/gf/v2/errors/gerror"
)

// Handler 处理发布到主题的事件。
type Handler func(ctx context.Context, event any) error

// Bus 是进程内事件总线。
type Bus struct {
	mu       sync.RWMutex
	handlers map[string][]Handler
}

var defaultBus = New()

// New 创建一个独立的事件总线。
func New() *Bus {
	return &Bus{
		handlers: make(map[string][]Handler),
	}
}

// Subscribe 在默认事件总线上订阅主题。
func Subscribe(topic string, handler Handler) error {
	return defaultBus.Subscribe(topic, handler)
}

// Publish 在默认事件总线上发布事件，并等待所有处理函数执行完成。
func Publish(ctx context.Context, topic string, event any) error {
	return defaultBus.Publish(ctx, topic, event)
}

// PublishAsync 在默认事件总线上异步发布事件。
func PublishAsync(ctx context.Context, topic string, event any) {
	defaultBus.PublishAsync(ctx, topic, event)
}

// Clear 清空默认事件总线上的所有订阅，主要用于测试。
func Clear() {
	defaultBus.Clear()
}

// Subscribe 订阅指定主题。
func (bus *Bus) Subscribe(topic string, handler Handler) error {
	if topic == "" {
		return gerror.New("eventbus: topic cannot be empty")
	}
	if handler == nil {
		return gerror.New("eventbus: handler cannot be nil")
	}

	bus.mu.Lock()
	defer bus.mu.Unlock()

	bus.handlers[topic] = append(bus.handlers[topic], handler)
	return nil
}

// Publish 发布事件，并等待所有处理函数执行完成。
func (bus *Bus) Publish(ctx context.Context, topic string, event any) error {
	handlers := bus.handlersOf(topic)

	var err error
	for _, handler := range handlers {
		if handlerErr := handler(ctx, event); handlerErr != nil && err == nil {
			err = handlerErr
		}
	}

	return err
}

// PublishAsync 异步发布事件。
func (bus *Bus) PublishAsync(ctx context.Context, topic string, event any) {
	handlers := bus.handlersOf(topic)
	if len(handlers) == 0 {
		return
	}

	go func() {
		_ = bus.publishHandlers(ctx, event, handlers)
	}()
}

// Clear 清空所有订阅。
func (bus *Bus) Clear() {
	bus.mu.Lock()
	defer bus.mu.Unlock()

	bus.handlers = make(map[string][]Handler)
}

func (bus *Bus) handlersOf(topic string) []Handler {
	bus.mu.RLock()
	defer bus.mu.RUnlock()

	handlers := bus.handlers[topic]
	copied := make([]Handler, len(handlers))
	copy(copied, handlers)
	return copied
}

func (bus *Bus) publishHandlers(ctx context.Context, event any, handlers []Handler) error {
	var err error
	for _, handler := range handlers {
		if handlerErr := handler(ctx, event); handlerErr != nil && err == nil {
			err = handlerErr
		}
	}
	return err
}
