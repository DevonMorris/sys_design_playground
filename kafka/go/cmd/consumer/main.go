package main

import (
	// "bufio"
	"fmt"
	// "log"
	"os"
	"os/signal"
  "syscall"
	"github.com/Shopify/sarama"
)

func connectConsumer(brokersUrl []string) (sarama.Consumer, error) {
  config := sarama.NewConfig()
  config.Consumer.Return.Errors = true

  conn, err := sarama.NewConsumer(brokersUrl, config)
  if (err != nil) {
    return nil, err
  }

  return conn, nil
}

func main() {
  topic:= "kafka_test"

	brokersUrl := []string{"localhost:9092"}
  worker, err := connectConsumer(brokersUrl)

  if err != nil {
    panic(err)
  }
	// Calling ConsumePartition. It will open one connection per broker
	// and share it for all partitions that live on it.
	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
  fmt.Println("Consumer Started")
  sigchan := make(chan os.Signal, 1)
  signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

  msgCount := 0

  doneCh := make(chan struct{})
  go func() {
    for {
      select {
        case err := <-consumer.Errors():
          fmt.Println(err)
        case msg := <-consumer.Messages():
          msgCount++
          fmt.Printf("Received msg count %d: | Topic(%s) | Message(%s)", msgCount, string(msg.Topic), string(msg.Value))
        case <-sigchan:
          fmt.Println("Interrupt detected")
          doneCh <- struct{}{}
      }
    }
  }()

  <-doneCh
  fmt.Println("Processed", msgCount, "messages")

  if err := worker.Close; err != nil {
    panic(err)
  }
}
