package application

import (
	"github.com/Andresito126/go-estudiantes/src/students/domain/entities"
	"github.com/Andresito126/go-estudiantes/src/students/domain/ports"
)

type FindAllStudentUseCase struct{
	studentRepository ports.IStudentRepository
}


func NewFindAllStudentUseCase(studentRepository ports.IStudentRepository) *FindAllStudentUseCase{
	return &FindAllStudentUseCase{studentRepository: studentRepository}
}

func (useCase *FindAllStudentUseCase) Run()([]entities.Student, error){
	return useCase.studentRepository.FindAll()
}


