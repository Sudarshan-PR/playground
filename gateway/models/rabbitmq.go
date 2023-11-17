package models

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

var RabbitMQClient *amqp.Connection
 
func CreateRabbitMQConnection(url string) (*amqp.Connection, error) {
	connection, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to rabbitmq.")
	return connection, nil
}

// PushToQueue pushes messages to a queue. Expects binding key, exchange and the message as parameters. Return error.
// Note: "body" parameter must be converted to []byte type before passing
func PushToQueue(binding_key string, exchange string, body []byte) error {
	//creating a channel
	ch, err := RabbitMQClient.Channel()
	if err != nil {
		fmt.Println("Error getting RMQ Channel: ", err)
		return err
	}
	defer ch.Close()

	//Pushing into the Queue based on binding_key
	err = ch.Publish(
		exchange,
		binding_key,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body: body,
		},
	)
	if err != nil {
		fmt.Println("Error getting RMQ Channel: ", err)
		return err
	}

	return nil
}
