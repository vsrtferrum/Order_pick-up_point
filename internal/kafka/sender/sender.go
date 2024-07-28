package sender

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/IBM/sarama"
	"gitlab.ozon.dev/berkinv/homework/internal/kafka/producer"
	"gitlab.ozon.dev/berkinv/homework/internal/models"
)

type KafkaSender struct {
	producer *producer.Producer
	topic    string
}
type Interface interface {
	SendMessage(message *models.LogMessage) error
	CreateMessage(command string, args []string) *models.LogMessage
}

func NewKafkaSender(producer *producer.Producer, topic string) *KafkaSender {
	return &KafkaSender{
		producer,
		topic,
	}
}

func (s *KafkaSender) SendMessage(message *models.LogMessage) error {
	kafkaMsg, err := s.buildMessage(*message)
	if err != nil {
		return err
	}

	_, _, err = s.producer.SendSyncMessage(kafkaMsg)
	if err != nil {
		return err
	}

	return nil
}

func (s *KafkaSender) CreateMessage(command string, args []string) *models.LogMessage {
	return &models.LogMessage{CreatedAt: time.Now(), CommandName: command, Args: args}
}

func (s *KafkaSender) buildMessage(message models.LogMessage) (*sarama.ProducerMessage, error) {
	msg, err := json.Marshal(message)

	if err != nil {
		return nil, err
	}

	return &sarama.ProducerMessage{
		Topic:     s.topic,
		Value:     sarama.ByteEncoder(msg),
		Partition: -1,
		Key:       sarama.StringEncoder(fmt.Sprint(message.AnswerID)),
	}, nil
}
