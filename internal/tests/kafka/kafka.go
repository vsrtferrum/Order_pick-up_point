//go:build integration

package kafkaArgs

import (
	"errors"
	"os"
)

type Implementation struct {
	Broker string
	Topic  string
}

func NewImplementation() (Implementation, error) {
	const brokerKey string = "KAFKA"
	const topicKey string = "TOPICNAME"

	broker := os.Getenv(brokerKey)
	if broker == "" {
		return Implementation{}, errors.New("Топик не иницилизирован")
	}

	topic := os.Getenv(topicKey)
	if topic == "" {
		return Implementation{}, errors.New("Топик не иницилизирован")
	}

	return Implementation{
		Broker: broker,
		Topic:  topic,
	}, nil
}
