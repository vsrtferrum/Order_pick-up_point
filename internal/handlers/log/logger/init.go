package log

import (
	"encoding/json"

	"github.com/IBM/sarama"
	"gitlab.ozon.dev/berkinv/homework/internal/errors"
	"gitlab.ozon.dev/berkinv/homework/internal/kafka/consumer"
	"gitlab.ozon.dev/berkinv/homework/internal/kafka/producer"
	"gitlab.ozon.dev/berkinv/homework/internal/kafka/reciever"
	"gitlab.ozon.dev/berkinv/homework/internal/kafka/sender"
	"gitlab.ozon.dev/berkinv/homework/internal/models"
)

func NewLogger(brokers []string, topic string) (*Log, error) {
	prod, err := producer.NewSyncProducer(brokers)
	if err != nil {
		return nil, err
	}

	cons, err := consumer.NewConsumer(brokers)
	if err != nil {
		return nil, err
	}

	send := sender.NewKafkaSender(prod, topic)

	handlers := map[string]reciever.HandleFunc{
		topic: func(message *sarama.ConsumerMessage) error {
			pm := models.LogMessage{}

			err = json.Unmarshal(message.Value, &pm)
			if err != nil {
				return errors.UnmarshalJsonErr
			}

			return nil
		},
	}
	rec := reciever.NewReceiver(cons, handlers)

	err = rec.Subscribe(topic)
	if err != nil {
		return nil, err
	}

	return &Log{
		reciever: rec,
		sender:   send,
	}, nil
}
