package standard

import (
	"context"
)

type ILogger interface {
	Debugc(ctx context.Context, msg string, values ...interface{})
	Infoc(ctx context.Context, msg string, values ...interface{})
	Warnc(ctx context.Context, msg string, values ...interface{})
	Errorc(ctx context.Context, msg string, values ...interface{})

	Debug(msg string, values ...interface{})
	Info(msg string, values ...interface{})
	Warn(msg string, values ...interface{})
	Error(msg string, values ...interface{})
}
