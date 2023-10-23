package main

import (
	"fmt"
	amqp "github.com/streadway/amqp"
	"time"
)

func main() {
	fmt.Println("RabbitMQ in golang...")
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		panic(err)
	}
	defer connection.Close()

	fmt.Println("Successfully connected to RabbitMQ instance")

	// producer
	channel, err := connection.Channel()

	if err != nil {
		panic(err)
	}

	// declaring queue with its properties over the the channel opened
	queue, err := channel.QueueDeclare(
		"testing", // name
		false,     // durable
		false,     // auto delete
		false,     // exclusive
		false,     // no wait
		nil,       // args
	)
	if err != nil {
		panic(err)
	}

	msg := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		Timestamp:    time.Now(),
		ContentType:  "text/plain",
		Body:         []byte("Go Go AMQP!"),
	}
	err = channel.Publish(
		"",
		"testing",
		false,
		false,
		msg,
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Queue status:", queue)
	fmt.Println("Successfully published message")

	msgs, err := channel.Consume(
		"testing",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)

	}

	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			fmt.Println("Received Message", msg.Body)
		}
	}()

	fmt.Println("Waiting for messages")
	<-forever
}
