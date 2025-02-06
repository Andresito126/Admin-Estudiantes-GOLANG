package ports

import "github.com/Andresito126/go-estudiantes/src/students/domain/entities"

type IStudentRepository interface {
	Save(student entities.Student) error
	FindAll()([]entities.Student, error)
	Update(id int, student entities.Student)error
	Delete(id int) error
	FindById(id int)([]entities.Student, error)
}
