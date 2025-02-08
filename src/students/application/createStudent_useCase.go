package application

import (
	"github.com/Andresito126/go-estudiantes/src/students/domain/ports"
	"github.com/Andresito126/go-estudiantes/src/students/domain/entities"
)

type CreateStudentUseCase struct {
	studentRepository ports.IStudentRepository
}

func NewCreateStudentUseCase(studentRepository ports.IStudentRepository) *CreateStudentUseCase {

	return &CreateStudentUseCase{studentRepository: studentRepository }
	
}

func (useCase *CreateStudentUseCase) Run(student entities.Student ) error {

	return useCase.studentRepository.Save(student)
}