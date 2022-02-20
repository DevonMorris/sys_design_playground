package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
  reader := bufio.NewReader(os.Stdin)
  fmt.Println("What message should we send?")
  payload, _ := reader.ReadString('\n')

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open channel on RabbiqMQ")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"golang-queue",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a Queue")

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(payload),
		},
	)

	failOnError(err, "Failed to publish message")
	log.Printf(" [x] Congrats, sending message: %s", payload)
}
