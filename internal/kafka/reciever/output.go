package reciever

import (
	"fmt"

	"github.com/IBM/sarama"
)

func output(pc sarama.PartitionConsumer, handler HandleFunc, partition int32, topic string) {
	for message := range pc.Messages() {
		handler(message)
		fmt.Println("Топик:", topic, "Раздел: ", partition, "Смещение:", message.Offset)
		fmt.Println("Полученный ключ:", string(message.Key), "Сообщение:", string(message.Value))
	}
}
