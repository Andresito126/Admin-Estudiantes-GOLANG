package application

import (
	"github.com/Andresito126/go-estudiantes/src/students/domain/entities"
	"github.com/Andresito126/go-estudiantes/src/students/domain/ports"
)

type UpdateStudentUseCase struct{
	db  ports.IStudentRepository
}

func NewUpdateStudentUseCase(db ports.IStudentRepository) *UpdateStudentUseCase{
	return &UpdateStudentUseCase{db : db}
	
}


func (uc *UpdateStudentUseCase) Run(id int, student entities.Student) error{
	return uc.db.Update(id, student)
	
}