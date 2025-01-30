package ports

import "github.com/Andresito126/go-estudiantes/src/students/domain/entities"

type IStudentRepository interface {
	Save(student entities.Student) error
	// FindAll()
	// Delete()
	// Edit()
}