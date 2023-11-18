package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

type queueBody struct {
	ID string `json:"client_id"`
	Code string `json:"code"`
}

func main() {
	// Load env variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client, err := createRabbitMQConnection(os.Getenv("RABBITMQ_ADDRESS"))
	if err != nil {
		log.Println("Error while creating connection to RabbitMQ:", err)
		return
	}
	ch, err := client.Channel()
	if err != nil {
		log.Println("Error while creating channel:", err)
		return
	}
	defer ch.Close()

	// Create Queue
	q, err := ch.QueueDeclare(
		"go",		// name
		true,		// durable
		false,		// delete when unused
		false,		// exclusive
		false,		// no-wait
		nil,		// arguments
	)

	if err = ch.QueueBind(q.Name, q.Name, "amq.direct", false, nil); err != nil {
		fmt.Println("Error binding RMQ Channel: ", err)
		return
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"Go Lister Service",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		fmt.Println("Error consuming RMQ Channel: ", err)
		return
	}

	forever := make(chan bool)

	go listener(msgs)
	fmt.Println("Listening to Queue: ", q.Name)
	<- forever
}

func createRabbitMQConnection(url string) (*amqp.Connection, error) {
	connection, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to rabbitmq.")
	return connection, nil
}

func listener(deliveries <-chan amqp.Delivery) {
	var (
		body queueBody
		output string
		err error
		success bool
	)

	for data := range(deliveries) {
		if err = json.Unmarshal(data.Body, &body); err != nil {
			log.Println("Invalid Queue Data: ", err)
			continue
		}

		output, err = runCode(body.Code) 
		success = true
		if err != nil {
			success = false
		}
		fmt.Println("Output: \n", output)
		fmt.Println("Success: ", success)
	}
}

func runCode(code string) (string, error) {
	filename := "/tmp/code-playground.go"
	err := ioutil.WriteFile(filename, []byte(code), 0644)
	if err != nil {
		return "", err
	}
	//defer os.Remove(filename)

	cmd := exec.Command("go", "run", filename)

	stdout, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error running code: ", err.Error())
	}

	return string(stdout), nil
}
