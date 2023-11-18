package setup

import (
	"log"
	"os"

	"github.com/Sudarshan-PR/playground/gateway/models"
	"github.com/joho/godotenv"
)

func Setup() error {
	// Load env variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mq, err := models.CreateRabbitMQConnection(os.Getenv("RABBITMQ_ADDRESS"))
	if err != nil {
		return err
	}
	models.RabbitMQClient = mq

	return nil
}
