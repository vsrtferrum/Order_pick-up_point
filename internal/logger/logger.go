package logger

import (
	"context"
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
)

var (
	global *slog.Logger

	lvlByEnv = map[string]slog.Level{
		"debug": slog.LevelDebug,
		"info":  slog.LevelInfo,
		"warn":  slog.LevelWarn,
		"error": slog.LevelError,
	}
)

func init() {
	lvlEnv := os.Getenv("LOG_LEVEL")

	lvl, ok := lvlByEnv[lvlEnv]
	if len(lvlEnv) == 0 && !ok {
		lvl = slog.LevelWarn
	}

	hConfig := &slog.HandlerOptions{
		Level: lvl,
	}

	var (
		writer io.Writer
		err    error
	)

	logWriter := os.Getenv("LOG_WRITER")
	switch logWriter {
	case "stdout":
		writer = os.Stdout
	case "stderr":
		writer = os.Stderr
	default:
		if len(logWriter) > 0 {
			var file *os.File

			file, err = os.OpenFile(logWriter, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
				log.Fatalf("Failed to open log writer file %q: %v", logWriter, err)
			}

			/*runtime.SetFinalizer(file, func() {
				ctx := context.Background()
				Warn(ctx, "closing file log writer")
				if err = file.Close(); err != nil {
					Errorf(ctx, "Failed to close log writer file %q: %v", logWriter, err)
				}
			})*/

			writer = io.MultiWriter(os.Stdout, file)
		} else {
			writer = os.Stdout
		}
	}

	var handler slog.Handler

	envHandler := os.Getenv("LOG_HANDLER")
	switch envHandler {
	case "text":
		handler = slog.NewTextHandler(writer, hConfig)
	case "json":
		handler = slog.NewJSONHandler(writer, hConfig)
	default:
		handler = slog.NewJSONHandler(writer, hConfig)
	}

	global = slog.New(handler)
}

func Debugf(ctx context.Context, format string, args ...interface{}) {
	Debug(ctx, fmt.Sprintf(format, args...))
}

func Debug(ctx context.Context, msg string, args ...interface{}) {
	l := logger(ctx)
	l.DebugContext(ctx, msg, args...)
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	Info(ctx, fmt.Sprintf(format, args...))
}

func Info(ctx context.Context, msg string, args ...interface{}) {
	l := logger(ctx)
	l.InfoContext(ctx, msg, args...)
}

func Warnf(ctx context.Context, format string, args ...interface{}) {
	Warn(ctx, fmt.Sprintf(format, args...))
}

func Warn(ctx context.Context, msg string, args ...interface{}) {
	l := logger(ctx)
	l.WarnContext(ctx, msg, args...)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	Error(ctx, fmt.Sprintf(format, args...))
}

func Error(ctx context.Context, msg string, args ...interface{}) {
	l := logger(ctx)
	l.ErrorContext(ctx, msg, args...)
}

type contextKey int

const (
	loggerContextKey contextKey = iota
)

func logger(ctx context.Context) *slog.Logger {
	l := global

	if ctxLogger, ok := ctx.Value(loggerContextKey).(*slog.Logger); ok {
		l = ctxLogger
	}

	return l
}
