package middlewares

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mattn/go-isatty"
	"github.com/shipengqi/log"
)

// Logger instances a Logger middleware that will write the logs to gin.DefaultWriter.
// By default, gin.DefaultWriter = os.Stdout.
func Logger() gin.HandlerFunc {
	return LoggerWithConfig(GetLoggerConfig(nil, nil, nil))
}

// GetLoggerConfig return gin.LoggerConfig which will write the logs to specified io.Writer with given gin.LogFormatter.
// By default, gin.DefaultWriter = os.Stdout
// reference: https://github.com/gin-gonic/gin#custom-log-format
func GetLoggerConfig(formatter gin.LogFormatter, output io.Writer, skipPaths []string) gin.LoggerConfig {
	if formatter == nil {
		formatter = GetDefaultLogFormatter()
	}

	return gin.LoggerConfig{
		Formatter: formatter,
		Output:    output,
		SkipPaths: skipPaths,
	}
}

// GetDefaultLogFormatter returns gin.LogFormatter.
func GetDefaultLogFormatter() gin.LogFormatter {
	return func(param gin.LogFormatterParams) string {
		var statusColor, methodColor, resetColor string
		if param.IsOutputColor() {
			statusColor = param.StatusCodeColor()
			methodColor = param.MethodColor()
			resetColor = param.ResetColor()
		}

		if param.Latency > time.Minute {
			// Truncate in a golang < 1.8 safe way
			param.Latency -= param.Latency % time.Second
		}

		return fmt.Sprintf("%s%3d%s - [%s] \"%v %s%s%s %s\" %s",
			// param.TimeStamp.Format("2006/01/02 - 15:04:05"),
			statusColor, param.StatusCode, resetColor,
			param.ClientIP,
			param.Latency,
			methodColor, param.Method, resetColor,
			param.Path,
			param.ErrorMessage,
		)
	}
}

func GetLoggerWithCtx(ctx context.Context) *log.Logger {
	lg := log.L()

	if requestID := ctx.Value(KeyRequestID); requestID != nil {
		lg = lg.WithValues(log.Any(KeyRequestID, requestID))
	}
	if username := ctx.Value(UsernameKey); username != nil {
		lg = lg.WithValues(log.Any(UsernameKey, username))
	}

	return lg
}

// LoggerWithConfig instance a Logger middleware with config.
func LoggerWithConfig(c gin.LoggerConfig) gin.HandlerFunc {
	formatter := c.Formatter
	if formatter == nil {
		formatter = GetDefaultLogFormatter()
	}

	out := c.Output
	if out == nil {
		out = gin.DefaultWriter
	}

	unlogged := c.SkipPaths

	isTerm := true

	if w, ok := out.(*os.File); !ok || os.Getenv("TERM") == "dumb" ||
		(!isatty.IsTerminal(w.Fd()) && !isatty.IsCygwinTerminal(w.Fd())) {
		isTerm = false
	}

	if isTerm {
		gin.ForceConsoleColor()
	}

	var skip map[string]struct{}

	if length := len(unlogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, path := range unlogged {
			skip[path] = struct{}{}
		}
	}

	return func(ctx *gin.Context) {
		// Start timer
		start := time.Now()
		path := ctx.Request.URL.Path
		raw := ctx.Request.URL.RawQuery

		// Process request
		ctx.Next()

		// Log only when path is not be skipped
		if _, ok := skip[path]; !ok {
			param := gin.LogFormatterParams{
				Request: ctx.Request,
				Keys:    ctx.Keys,
			}

			// Stop timer
			param.TimeStamp = time.Now()
			param.Latency = param.TimeStamp.Sub(start)

			param.ClientIP = ctx.ClientIP()
			param.Method = ctx.Request.Method
			param.StatusCode = ctx.Writer.Status()
			param.ErrorMessage = ctx.Errors.ByType(gin.ErrorTypePrivate).String()
			param.BodySize = ctx.Writer.Size()

			if raw != "" {
				path = path + "?" + raw
			}

			param.Path = path

			GetLoggerWithCtx(ctx).Infot(formatter(param))
		}
	}
}
