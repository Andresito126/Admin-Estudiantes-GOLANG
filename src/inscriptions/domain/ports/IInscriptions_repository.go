package ports

import "github.com/Andresito126/go-estudiantes/src/inscriptions/domain/entities"

type IInscriptionsRepository interface {

	Save(inscription *entities.Inscription) error

	FindAll() ([]entities.Inscription, error)
	
}