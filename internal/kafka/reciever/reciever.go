//go:generate minimock -i gitlab.ozon.dev/berkinv/homework/internal/kafka/reciever -o ./mocks/reciever_mock.go -n reciever_mock
package reciever

import (
	"github.com/IBM/sarama"
	"gitlab.ozon.dev/berkinv/homework/internal/errors"
	"gitlab.ozon.dev/berkinv/homework/internal/kafka/consumer"
)

type HandleFunc func(message *sarama.ConsumerMessage) error

type KafkaReceiver struct {
	Consumer *consumer.Consumer
	Handlers map[string]HandleFunc
}

type Interface interface {
	Subscribe(topic string) error
}

func NewReceiver(consumer *consumer.Consumer, handlers map[string]HandleFunc) *KafkaReceiver {
	return &KafkaReceiver{
		Consumer: consumer,
		Handlers: handlers,
	}
}

func (r *KafkaReceiver) Subscribe(topic string) error {
	handler, ok := r.Handlers[topic]

	if !ok {
		return errors.SubscribeErr
	}

	partitionList, err := r.Consumer.Partitions(topic)

	if err != nil {
		return err
	}

	initialOffset := sarama.OffsetNewest

	for _, partition := range partitionList {
		pc, err := r.Consumer.ConsumePartition(topic, partition, initialOffset)

		if err != nil {
			return err
		}

		go func(pc sarama.PartitionConsumer, partition int32) {
			output(pc, handler, partition, topic)
		}(pc, partition)
	}
	return nil
}
