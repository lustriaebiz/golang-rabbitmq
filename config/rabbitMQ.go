package config

import (
	"fmt"

	"github.com/streadway/amqp"
)

func RabbitMQ() (*amqp.Channel, error) {

	conn, err := amqp.Dial("amqp://guest:1234@localhost:5672/")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	ch, err := conn.Channel()

	if err != nil {
		fmt.Println(err)
	}

	return ch, nil

}
