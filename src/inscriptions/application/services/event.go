package services

import (
    "github.com/Andresito126/go-estudiantes/src/inscriptions/application/repository"
    "github.com/Andresito126/go-estudiantes/src/inscriptions/domain/entities"
)

type EventService struct {
    rabbitRepo repository.IRabbitRepository
}

func NewEventService(rabbitRepo repository.IRabbitRepository) *EventService {
    return &EventService{
        rabbitRepo: rabbitRepo,
    }
}

func (e *EventService) PublishEvent(inscription entities.Inscription) error {
    return e.rabbitRepo.Publish(inscription)
}