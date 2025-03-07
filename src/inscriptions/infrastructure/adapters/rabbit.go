package adapters

import (
    "encoding/json"
    "log"

    "github.com/Andresito126/go-estudiantes/src/inscriptions/domain/entities"
    "github.com/streadway/amqp"
)

type RabbitRepository struct {
    conn *amqp.Connection
}

func NewRabbitRepository(url string) *RabbitRepository {
    conn, err := amqp.Dial(url)
    if err != nil {
        log.Fatalf("No se pudo conectar a RabbitMQ: %v", err)
    }
    return &RabbitRepository{
        conn: conn,
    }
}

func (r *RabbitRepository) Publish(inscription entities.Inscription) error {
    ch, err := r.conn.Channel()
    if err != nil {
        return err
    }
    defer ch.Close()

    q, err := ch.QueueDeclare(
        "inscriptionsQueue", // nombre de la cola
        true    ,               // durable
        false,               // delete when unused
        false,               // exclusive
        false,               // no-wait
        nil,                 // arguments
    )
    if err != nil {
        return err
    }

    // Convertir la inscripción a JSON
    body, err := json.Marshal(inscription)
    if err != nil {
        return err
    }

    err = ch.Publish(
        "",     // exchange
        q.Name, // routing key
        false,  // mandatory
        false,  // immediate
        amqp.Publishing{
            ContentType: "application/json",
            Body:        body,
        },
    )
    if err != nil {
        return err
    }

    log.Printf("Mensaje enviado a la cola %s: %s", q.Name, string(body))
    return nil
}