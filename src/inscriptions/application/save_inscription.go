package application

import (
	"github.com/Andresito126/go-estudiantes/src/inscriptions/domain/entities"
	"github.com/Andresito126/go-estudiantes/src/inscriptions/domain/ports"
)

type SaveInscriptionUseCase struct {

	inscriptionRepository ports.IInscriptionsRepository
}

func NewSaveInscriptionUseCase(inscriptionRepository ports.IInscriptionsRepository)*SaveInscriptionUseCase{

	return &SaveInscriptionUseCase{inscriptionRepository: inscriptionRepository}

}


func (useCase *SaveInscriptionUseCase) Run(inscription *entities.Inscription) error{

	return useCase.inscriptionRepository.Save(inscription)
}