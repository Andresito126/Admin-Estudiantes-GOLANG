package ports

import "github.com/Andresito126/go-estudiantes/src/courses/domain/entities"

type ICourseRepository interface {
	Save(course entities.Course) error
	FindAll()([] entities.Course, error)
	Update(id int, course entities.Course) error
	Delete(id int ) error
}