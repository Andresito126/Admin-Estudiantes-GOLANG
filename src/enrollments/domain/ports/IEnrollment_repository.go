package ports

import "github.com/Andresito126/go-estudiantes/src/enrollments/domain/entities"

type IEnrollmentRepository interface {
	CreateEnrollment(enrollment *entities.Enrollment) error
	// FindEnrollmentsByCourseID(courseID int) ([]entities.Enrollment, error)
	// DeleteEnrollment(enrollmentID int) error
}
