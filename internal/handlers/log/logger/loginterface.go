package log

import (
	"context"

	"gitlab.ozon.dev/berkinv/homework/internal/kafka/reciever"
	"gitlab.ozon.dev/berkinv/homework/internal/kafka/sender"
)

type LogInterface interface {
	Input(command string, arg []string) error
	Error(ctx context.Context, err error) error
}
type Log struct {
	LogInterface
	sender   *sender.KafkaSender
	reciever *reciever.KafkaReceiver
}
