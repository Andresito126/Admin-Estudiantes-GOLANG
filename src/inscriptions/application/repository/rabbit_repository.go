package repository

import "github.com/Andresito126/go-estudiantes/src/inscriptions/domain/entities"

type IRabbitRepository interface {
    Publish(inscription entities.Inscription) error
}