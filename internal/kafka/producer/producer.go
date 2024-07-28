package producer

import (
	"fmt"
	"time"

	"github.com/IBM/sarama"
	"gitlab.ozon.dev/berkinv/homework/internal/errors"
)

type Producer struct {
	brokers      []string
	syncProducer sarama.SyncProducer
}

func NewSyncProducer(brokers []string) (*Producer, error) {
	syncProducerConfig := sarama.NewConfig()
	const (
		producerTimeout time.Duration = 3 * time.Second
		maxMessages     int           = 100
	)
	syncProducerConfig.Producer.Partitioner = sarama.NewHashPartitioner
	syncProducerConfig.Producer.RequiredAcks = sarama.WaitForLocal
	syncProducerConfig.Producer.Timeout = producerTimeout
	syncProducerConfig.Producer.Return.Successes = true
	syncProducerConfig.Producer.Return.Errors = true
	syncProducerConfig.Producer.Flush.MaxMessages = maxMessages

	syncProducer, err := sarama.NewSyncProducer(brokers, syncProducerConfig)

	if err != nil {
		return nil, errors.CreationSyncKafkaProduserErr
	}

	return &Producer{syncProducer: syncProducer}, nil
}

func (k *Producer) SendSyncMessage(message *sarama.ProducerMessage) (partition int32, offset int64, err error) {
	return k.syncProducer.SendMessage(message)
}

func (k *Producer) SendSyncMessages(messages []*sarama.ProducerMessage) error {
	err := k.syncProducer.SendMessages(messages)
	if err != nil {
		fmt.Println("kafka.Connector.SendMessages error", err)
	}

	return err
}
func (producer Producer) Close() error {
	return producer.syncProducer.Close()
}
