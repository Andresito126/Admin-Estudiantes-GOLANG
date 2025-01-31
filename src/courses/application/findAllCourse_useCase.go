package application

import (
	"github.com/Andresito126/go-estudiantes/src/courses/domain/entities"
	"github.com/Andresito126/go-estudiantes/src/courses/domain/ports"
)

type FindAllCourseUseCase struct {
	courseRepository ports.ICourseRepository
}

func NewFindAllCourseUseCase(courseRepository ports.ICourseRepository)*FindAllCourseUseCase{
	return &FindAllCourseUseCase{courseRepository: courseRepository}
}

func(useCase *FindAllCourseUseCase)Run()([]entities.Course, error){
	return useCase.courseRepository.FindAll()
}

