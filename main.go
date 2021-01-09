package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	fmt.Println("Rabbit MQ test")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	log.Println("Successfuly connect to RebbitMQ(localhost:5672)")
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(q)

	err = ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte("Hello world")},
	)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfuly published message to a queue")
}
