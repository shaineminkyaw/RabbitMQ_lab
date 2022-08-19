package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {

	//connect to rabbitmq instance from go code
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()
	fmt.Println("Hello from RabbitMQ...")
	fmt.Println("Succesfully Connected To our RabbitMQ instance...")

	//obtain a connection create message queue channel on RabbitMQ instance
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer ch.Close()

	//declare new queue from RabbitMQ instance
	queue, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(queue)

	//channel launch to publish
	err = ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello Wolrd..."),
		},
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Successfully Published Message to Queue")

}
