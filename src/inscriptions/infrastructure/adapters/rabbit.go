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

func NewRabbitRepository(conn *amqp.Connection) *RabbitRepository {
    return &RabbitRepository{conn: conn}
}

func (r *RabbitRepository) Publish(inscription entities.Inscription) error {
    ch, err := r.conn.Channel()
    if err != nil {
        return err
    }
    defer ch.Close()

    // declara a el exchange inscriptions_exchange
    err = ch.ExchangeDeclare(
        "inscriptions_exchange", // name
        "direct",                // Tipo de exchange 
        true,                    // Durable
        false,                   // Auto-deleted
        false,                   // Internal
        false,                   // No-wait
        nil,                     // Arguments
    )
    if err != nil {
        return err
    }

    //  cola de inscripciones
    q, err := ch.QueueDeclare(
        "inscriptionsQueue", // name
        true,                // Durable
        false,               // Delete when unused
        false,               // Exclusive
        false,               // No-wait
        nil,                 // Arguments
    )
    if err != nil {
        return err
    }

    // vicila la cola al exchange usando la routing key
    err = ch.QueueBind(
        q.Name,                     // Queue name
        "inscriptions_routing_key", // Routing key
        "inscriptions_exchange",    // Exchange name
        false,                      // No-wait
        nil,                        // Arguments
    )
    if err != nil {
        return err
    }

    // convierte a json
    body, err := json.Marshal(inscription)
    if err != nil {
        return err
    }

    // publica el mensaje en el exchange con la routing key
    err = ch.Publish(
        "inscriptions_exchange",    // name del exchange
        "inscriptions_routing_key", // Routing key
        false,                      // Mandatory
        false,                      // Immediate
        amqp.Publishing{
            ContentType: "application/json",
            Body:        body,
        },
    )
    if err != nil {
        return err
    }

    log.Printf("Mensaje enviado al exchange 'inscriptions_exchange' con routing key 'inscriptions_routing_key' hacia la cola %s: %s", q.Name, string(body))
    return nil
}
