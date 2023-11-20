package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"

	pb "github.com/Sudarshan-PR/playground/go-playground/notification-protos"
	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var conn *grpc.ClientConn

type queueBody struct {
	ID string `json:"client_id"`
	Code string `json:"code"`
	UserID string `json:"userid"`
}

type message struct {
	Delivery <-chan amqp.Delivery
	Close func() error
}

func main() {
	// Load env variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	client, err := createRabbitMQConnection(os.Getenv("RABBITMQ_ADDRESS"))
	if err != nil {
		log.Println("Error while creating connection to RabbitMQ:", err)
		return
	}
	msgs, err := setupRabbitMQChannel(client)
	if err != nil {
		log.Println("Error while creating listener:", err)
		return
	}
	defer msgs.Close()

	// Setup gRPC Client
	conn, err = grpc.Dial(os.Getenv("NOTIFICATIONS_SERVICE_ADDRESS"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to notification service: %v", err)
	}
	defer conn.Close()
	fmt.Println("Successfully connected to Notification service")

	forever := make(chan bool)

	go listener(msgs.Delivery)

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

func setupRabbitMQChannel(client *amqp.Connection) (message, error) {
	ch, err := client.Channel()
	if err != nil {
		log.Println("Error while creating channel:", err)
		return message{}, err
	}
	//defer ch.Close()

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
		return message{}, err
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
		return message{}, err
	}

	fmt.Println("Listening to Queue: ", q.Name)
	return message{Delivery: msgs, Close: ch.Close}, err
}

func listener(deliveries <-chan amqp.Delivery) {
	var (
		body queueBody
		output string
		outputType string
		err error
	)

	for data := range(deliveries) {
		if err = json.Unmarshal(data.Body, &body); err != nil {
			log.Println("Invalid Queue Data: ", err)
			continue
		}

		fmt.Println("Message received: ", body)
		output, err = runCode(body.Code) 
		outputType = "success"
		if err != nil {
			outputType = "failed"
		}
		fmt.Println("Output: \n", output)
		fmt.Println("Success: ", outputType)
		
		err = sendOutputToNotification(body.UserID, output, outputType)
		if err != nil {
			log.Println("Could not send to notification service: ", err)
		}
	}
}

func sendOutputToNotification(userid, output, outputType string) error {
	c := pb.NewNotificationClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := c.PushNotification(ctx, &pb.NotificationRequest{Output: output, Userid: userid, Type: outputType})
	if err != nil {
		return err
	}
	fmt.Println(resp)
	return nil
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
