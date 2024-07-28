package logger

import (
	log "gitlab.ozon.dev/berkinv/homework/internal/handlers/log/logger"
	"gitlab.ozon.dev/berkinv/homework/internal/handlers/log/nonlog"
)

func NewLogger(swap bool, brokers []string, topic string) *Logger {
	var (
		loger     *log.Log
		nonlog    *nonlog.Nonlog
		loggerErr error
	)
	if swap {
		loger, loggerErr = log.NewLogger(brokers, topic)
		if loggerErr != nil {
			return nil
			panic(loggerErr)
		}
	}
	return &Logger{
		swap:    swap,
		withlog: loger,
		nonlog:  nonlog,
	}
}
