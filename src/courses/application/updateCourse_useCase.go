package application

import (
	"github.com/Andresito126/go-estudiantes/src/courses/domain/entities"
	"github.com/Andresito126/go-estudiantes/src/courses/domain/ports"
)

type UpdateCourseUseCase struct {
	courseRepository ports.ICourseRepository
}

func NewUpdateCourseUseCase(courseRepository ports.ICourseRepository)*UpdateCourseUseCase{
	return &UpdateCourseUseCase{courseRepository: courseRepository}
}

func(uc *UpdateCourseUseCase)Run(id int, course entities.Course)error{
	return uc.courseRepository.Update(id, course)
}