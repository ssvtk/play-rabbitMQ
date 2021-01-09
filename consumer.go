package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	fmt.Println("RabbitMQ Consumer app")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()
	msgs, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	forever := make(chan bool)
	go func() {
for d := range msgs {
	fmt.Printf("Received message %s\n", d.Body)
}
	}()

	fmt.Println("Seccesfully connected to our RabbitMQ instance")
	fmt.Println("[*] - waiting for messages")
	<-forever
}
