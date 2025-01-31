package application

import "github.com/Andresito126/go-estudiantes/src/courses/domain/ports"

type DeleteCourseUseCase struct {
	courseRepository ports.ICourseRepository
}

func NewDeleteCourseUseCase(courseRepository ports.ICourseRepository)*DeleteCourseUseCase{
	return &DeleteCourseUseCase{courseRepository:courseRepository}
}

func(uc *DeleteCourseUseCase)Run(id int) error{
	return uc.courseRepository.Delete(id)
}