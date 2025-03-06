package repository

type IRabbitRepository interface {
    Publish(queue string, message string) error
}