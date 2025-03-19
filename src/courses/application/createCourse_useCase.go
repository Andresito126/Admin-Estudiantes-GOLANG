package application

import (
	"github.com/Andresito126/go-estudiantes/src/courses/domain/entities"
	"github.com/Andresito126/go-estudiantes/src/courses/domain/ports"
)

type SaveCourseUseCase struct {
    courseRepository ports.ICourseRepository
}

func NewSaveCourseUseCase(courseRepository ports.ICourseRepository) *SaveCourseUseCase {
  
    return &SaveCourseUseCase{courseRepository : courseRepository}
    
}

func (useCase *SaveCourseUseCase) Run(course entities.Course) (int, error) {
	id, err := useCase.courseRepository.Save(course)
	return id, err
}

