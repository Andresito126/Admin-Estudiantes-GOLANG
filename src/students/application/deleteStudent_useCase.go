package application

import (
	
	"github.com/Andresito126/go-estudiantes/src/students/domain/ports"
)

type DeleteStudentUseCase struct {
	db ports.IStudentRepository
}

func NewDeleteStudentUseCase(db ports.IStudentRepository) *DeleteStudentUseCase{
	return &DeleteStudentUseCase{db : db}
}

func( uc * DeleteStudentUseCase)Run(id int) error{
	return uc.db.Delete(id)
}
