package setup

import (
	"os"

	"github.com/Sudarshan-PR/playground/gateway/models"
)

func Setup() error {
	mq, err := models.CreateRabbitMQConnection(os.Getenv("RABBITMQ_ADDRESS"))
	if err != nil {
		return err
	}
	models.RabbitMQClient = mq

	return nil
}
