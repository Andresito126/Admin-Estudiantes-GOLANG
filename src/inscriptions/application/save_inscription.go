package application

import (
	"github.com/Andresito126/go-estudiantes/src/inscriptions/application/services"
	"github.com/Andresito126/go-estudiantes/src/inscriptions/domain/entities"
	"github.com/Andresito126/go-estudiantes/src/inscriptions/infrastructure/adapters"
)

type SaveInscriptionUseCase struct {
    inscriptionRepo *adapters.MySQL
    eventService    *services.EventService
}

func NewSaveInscriptionUseCase(inscriptionRepo *adapters.MySQL, eventService *services.EventService) *SaveInscriptionUseCase {
    return &SaveInscriptionUseCase{
        inscriptionRepo: inscriptionRepo,
        eventService:    eventService,
    }
}

func (useCase *SaveInscriptionUseCase) Run(inscription *entities.Inscription) error {
    
    err := useCase.inscriptionRepo.Save(inscription)
    if err != nil {
        return err
    }

    // publica el evento en rabit
    err = useCase.eventService.PublishEvent("inscriptionQueue", "Nueva inscripción solicitada")
    if err != nil {
        return err
    }

    return nil
}