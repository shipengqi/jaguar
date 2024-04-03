package logging

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/shipengqi/log"
)

// InterceptorLogger adapts custom logger to interceptor logger.
func InterceptorLogger(l *log.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, level logging.Level, msg string, fields ...any) {
		f := make([]log.Field, 0, len(fields)/2)

		for i := 0; i < len(fields); i += 2 {
			key := fields[i]
			value := fields[i+1]

			switch v := value.(type) {
			case string:
				f = append(f, log.String(key.(string), v))
			case int:
				f = append(f, log.Int(key.(string), v))
			case bool:
				f = append(f, log.Bool(key.(string), v))
			default:
				f = append(f, log.Any(key.(string), v))
			}
		}

		logger := l.WithValues(f...)

		switch level {
		case logging.LevelDebug:
			logger.Debug(msg)
		case logging.LevelInfo:
			logger.Info(msg)
		case logging.LevelWarn:
			logger.Warn(msg)
		case logging.LevelError:
			logger.Error(msg)
		default:
			logger.Debug(msg)
		}
	})
}
