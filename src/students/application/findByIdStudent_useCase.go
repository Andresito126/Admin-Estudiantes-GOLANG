package application

import (
	"github.com/Andresito126/go-estudiantes/src/students/domain/entities"
	"github.com/Andresito126/go-estudiantes/src/students/domain/ports"
)

type FindByIdUseCase struct{

	db ports.IStudentRepository
}

func NewFindByIdUseCase (db ports.IStudentRepository) * FindByIdUseCase{
	return &FindByIdUseCase{db:  db}
}

func (uc * FindByIdUseCase)Run (id int)( []entities.Student, error) {
	return uc.db.FindById(id)
}
