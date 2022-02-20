package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"github.com/Shopify/sarama"
)

func connectProducer(brokersUrl []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	conn, err := sarama.NewSyncProducer(brokersUrl, config)
	if err != nil {
		return nil, err
	}

	return conn, err
}

func pushCommentToTopic(topic string, message []byte) error {
	brokersUrl := []string{"localhost:9092"}
	producer, err := connectProducer(brokersUrl)
	if err != nil {
		return err
	}

	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := producer.SendMessage(msg)
  if err != nil {
    return err
  }

	fmt.Printf("Message is stored on topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
  return nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
  reader := bufio.NewReader(os.Stdin)
  fmt.Println("What message should we send?")
  payload, _ := reader.ReadString('\n')

  err := pushCommentToTopic("kafka_test", []byte(payload))
	failOnError(err, "Failed to publish message")
}
