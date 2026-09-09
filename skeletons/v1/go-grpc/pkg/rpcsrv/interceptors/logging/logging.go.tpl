package logging

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"go.uber.org/zap"
)

// InterceptorLogger adapts zap SugaredLogger to interceptor logger.
func InterceptorLogger(l *zap.SugaredLogger) logging.Logger {
	return logging.LoggerFunc(func(_ context.Context, level logging.Level, msg string, fields ...any) {
		args := make([]any, 0, len(fields))
		for i := 0; i+1 < len(fields); i += 2 {
			args = append(args, fields[i], fields[i+1])
		}
		switch level {
		case logging.LevelDebug:
			l.Debugw(msg, args...)
		case logging.LevelInfo:
			l.Infow(msg, args...)
		case logging.LevelWarn:
			l.Warnw(msg, args...)
		case logging.LevelError:
			l.Errorw(msg, args...)
		default:
			l.Debugw(msg, args...)
		}
	})
}
