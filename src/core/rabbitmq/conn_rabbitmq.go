package rabbitmq

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/streadway/amqp"
	"github.com/joho/godotenv"
)

var (
	instance *amqp.Connection
	once     sync.Once
)


func Connect() (*amqp.Connection, error) {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error al cargar el .env: %v", err)
		}

		user := os.Getenv("RABBITMQ_USER")
		password := os.Getenv("RABBITMQ_PASSWORD")
		host := os.Getenv("RABBITMQ_HOST")
		port := os.Getenv("RABBITMQ_PORT")

		rabbitURL := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, host, port)

		conn, err := amqp.Dial(rabbitURL)
		if err != nil {
			log.Fatalf("Error al conectar con RabbitMQ: %v", err)
		}

		instance = conn
		log.Println("Conexión a RabbitMQ exitosa")
	})

	return instance, nil
}
