package services

import (
    "log"

    "github.com/Andresito126/go-estudiantes/src/inscriptions/application/repository"
)

type EventService struct {
    rabbitRepo repository.IRabbitRepository
}

func NewEventService(rabbitRepo repository.IRabbitRepository) *EventService {
    return &EventService{
        rabbitRepo: rabbitRepo,
    }
}

func (s *EventService) PublishEvent(queue string, message string) error {
    if s.rabbitRepo == nil {
        log.Fatal("RabbitRepo no inicializado")
    }
    return s.rabbitRepo.Publish(queue, message)
}