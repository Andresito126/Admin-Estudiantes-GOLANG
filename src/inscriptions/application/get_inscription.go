package application

import (
	"github.com/Andresito126/go-estudiantes/src/inscriptions/domain/entities"
	"github.com/Andresito126/go-estudiantes/src/inscriptions/domain/ports"
)

type GetAllInscriptionsUseCase struct {
	inscriptionRepository ports.IInscriptionsRepository
}

func NewGetAllInscriptionsUseCase(inscriptionRepository ports.IInscriptionsRepository) *GetAllInscriptionsUseCase {
	return &GetAllInscriptionsUseCase{inscriptionRepository: inscriptionRepository}
}


func (useCase *GetAllInscriptionsUseCase) Run() ([]entities.Inscription, error) {
	return useCase.inscriptionRepository.FindAll()
}
