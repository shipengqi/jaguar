package xlog

import (
	"os"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	mu     sync.RWMutex
	global *zap.SugaredLogger
)

// Init initializes the global logger from Options.
func Init(opts *Options) {
	mu.Lock()
	defer mu.Unlock()
	global = newLogger(opts).Sugar()
}

// Close flushes any buffered log entries.
func Close() error {
	mu.RLock()
	defer mu.RUnlock()
	if global == nil {
		return nil
	}
	return global.Sync()
}

func newLogger(opts *Options) *zap.Logger {
	var level zapcore.Level
	_ = level.UnmarshalText([]byte(opts.Level))

	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeTime:     func(t time.Time, enc zapcore.PrimitiveArrayEncoder) { enc.AppendString(t.Format("2006-01-02 15:04:05.000")) },
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	if opts.Format == "json" {
		encoderCfg.EncodeLevel = zapcore.LowercaseLevelEncoder
	} else if opts.EnableColor {
		encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	}

	cores := make([]zapcore.Core, 0, len(opts.OutputPaths))
	for _, path := range opts.OutputPaths {
		var w zapcore.WriteSyncer
		if path == "stdout" {
			w = zapcore.AddSync(os.Stdout)
		} else if path == "stderr" {
			w = zapcore.AddSync(os.Stderr)
		} else {
			w = zapcore.AddSync(&lumberjack.Logger{
				Filename:   path,
				MaxSize:    100,
				MaxBackups: 5,
				MaxAge:     30,
				Compress:   true,
			})
		}

		enc := zapcore.NewConsoleEncoder(encoderCfg)
		if opts.Format == "json" {
			enc = zapcore.NewJSONEncoder(encoderCfg)
		}
		cores = append(cores, zapcore.NewCore(enc, w, level))
	}

	opts2 := []zap.Option{zap.AddCallerSkip(1)}
	if !opts.DisableCaller {
		opts2 = append(opts2, zap.WithCaller(true))
	}

	return zap.New(zapcore.NewTee(cores...), opts2...)
}

func std() *zap.SugaredLogger {
	mu.RLock()
	defer mu.RUnlock()
	if global == nil {
		l, _ := zap.NewProduction()
		return l.Sugar()
	}
	return global
}

// WithValues returns a logger with additional key-value pairs.
func WithValues(keysAndValues ...any) *zap.SugaredLogger {
	return std().With(keysAndValues...)
}

func Debug(args ...any)                             { std().Debug(args...) }
func Info(args ...any)                              { std().Info(args...) }
func Warn(args ...any)                              { std().Warn(args...) }
func Error(args ...any)                             { std().Error(args...) }
func Fatal(args ...any)                             { std().Fatal(args...) }
func Debugf(format string, args ...any)             { std().Debugf(format, args...) }
func Infof(format string, args ...any)              { std().Infof(format, args...) }
func Warnf(format string, args ...any)              { std().Warnf(format, args...) }
func Errorf(format string, args ...any)             { std().Errorf(format, args...) }
func Fatalf(format string, args ...any)             { std().Fatalf(format, args...) }
func Debugw(msg string, keysAndValues ...any)       { std().Debugw(msg, keysAndValues...) }
func Infow(msg string, keysAndValues ...any)        { std().Infow(msg, keysAndValues...) }
func Warnw(msg string, keysAndValues ...any)        { std().Warnw(msg, keysAndValues...) }
func Errorw(msg string, keysAndValues ...any)       { std().Errorw(msg, keysAndValues...) }
func Fatalw(msg string, keysAndValues ...any)       { std().Fatalw(msg, keysAndValues...) }

// Logger returns the global sugared logger.
func Logger() *zap.SugaredLogger { return std() }

// Sync flushes any buffered log entries.
func Sync() error { return Close() }
