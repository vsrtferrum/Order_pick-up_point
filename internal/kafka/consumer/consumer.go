package consumer

import (
	"time"

	"github.com/IBM/sarama"
)

type Consumer struct {
	brokers  []string
	Consumer sarama.Consumer
}

func NewConsumer(brokers []string) (*Consumer, error) {
	ConsumerConfig := sarama.NewConfig()

	const autoCommitInterval = 3 * time.Second

	ConsumerConfig.Consumer.Return.Errors = false
	ConsumerConfig.Consumer.Offsets.AutoCommit.Enable = true
	ConsumerConfig.Consumer.Offsets.AutoCommit.Interval = autoCommitInterval
	ConsumerConfig.Consumer.Offsets.Initial = sarama.OffsetNewest

	consumer, err := sarama.NewConsumer(brokers, ConsumerConfig)
	if err != nil {
		return nil, err
	}

	return &Consumer{
		brokers:  brokers,
		Consumer: consumer,
	}, nil
}

func (i Consumer) ConsumePartition(
	topic string,
	partition int32,
	offset int64,
) (sarama.PartitionConsumer, error) {
	return i.Consumer.ConsumePartition(topic, partition, offset)
}

func (i Consumer) Partitions(topic string) ([]int32, error) {
	return i.Consumer.Partitions(topic)
}

func (i Consumer) Close() error {
	return i.Consumer.Close()
}
