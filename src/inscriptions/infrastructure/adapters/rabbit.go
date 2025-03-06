package adapters

import (
    "log"

    "github.com/streadway/amqp"
)

type RabbitRepository struct {
    conn *amqp.Connection
}

func NewRabbitRepository(url string) *RabbitRepository {
    conn, err := amqp.Dial(url)
    if err != nil {
        log.Fatalf("no se pudo conectar a rabbit: %v", err)
    }
    return &RabbitRepository{
        conn: conn,
    }
}

func (r *RabbitRepository) Publish(inscriptionQueue string, message string) error {
    ch, err := r.conn.Channel()
    if err != nil {
        return err
    }
    defer ch.Close()

    q, err := ch.QueueDeclare(
        inscriptionQueue, // nombre de la cola
        true, // durable
        false, // delete when unused
        false, // exclusive
        false, // no-wait
        nil,   // arguments
    )
    if err != nil {
        return err
    }

    err = ch.Publish(
        "",     // exchange
        q.Name, // routing key
        false,  // mandatory
        false,  // immediate
        amqp.Publishing{
            ContentType: "text/plain",
            Body:        []byte(message),
        },
    )
    if err != nil {
        return err
    }

    log.Printf("Mensaje enviado a la cola %s: %s", inscriptionQueue, message)
    return nil
}