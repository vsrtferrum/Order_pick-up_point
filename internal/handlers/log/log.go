package logger

import (
	log "gitlab.ozon.dev/berkinv/homework/internal/handlers/log/logger"
	"gitlab.ozon.dev/berkinv/homework/internal/handlers/log/nonlog"
	"gitlab.ozon.dev/berkinv/homework/internal/models"
)

type swapInterface interface {
	Input(input models.LogMessage) error
}

type Logger struct {
	nonlog  *nonlog.Nonlog
	withlog *log.Log
	swapInterface
	swap bool
}

func (log *Logger) Input(command string, args []string) error {
	if log.swap {
		err := log.withlog.Input(command, args)
		return err
	} else {
		log.nonlog.Input(command, args)
	}
	return nil
}
