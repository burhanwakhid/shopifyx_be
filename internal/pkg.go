package internal

import (
	"context"
	"time"

	"github.com/burhanwakhid/shopifyx_backend/pkg/translator"
)

type Producer interface {
	PublishAsyncMessage(ctx context.Context, topic string, msg []byte, key string)
}

type Telemetry interface {
	IncrementCounter(bucket string, tags ...string)
	Gauge(bucket string, value interface{}, tags ...string)
	Histogram(bucket string, value interface{}, tags ...string)
	Flush()
}

type Locker interface {
	Lock(ctx context.Context, key string, ttl time.Duration) error
	ReleaseLock(ctx context.Context, keys ...string) error
	IsLocked(err error) bool
}

type Logger interface {
	Error(ctx context.Context, args ...interface{})
	ErrorWithFields(ctx context.Context, err error, msg string, fields map[string]interface{}, args ...interface{})
	Set(ctx context.Context, key string, value interface{})
}

type ErrorTranslator interface {
	Translate(locale, code string, vars ...interface{}) translator.ErrorTemplate
}
