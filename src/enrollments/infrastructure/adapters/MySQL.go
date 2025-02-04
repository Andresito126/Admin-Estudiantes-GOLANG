package adapters

import (
	"database/sql"

	"github.com/Andresito126/go-estudiantes/src/enrollments/domain/entities"
	
)

type MySQL struct{
	db *sql.DB
}

func NewMySQL (db *sql.DB) *MySQL{
	return &MySQL{db: db}
}

func (mysql *MySQL) CreateEnrollment(enrollment *entities.Enrollment)error{
	query := "INSERT INTO enrollments (student_id, course_id) VALUES (?, ?)"
	result, err := mysql.db.Exec(query, enrollment.StudentID, enrollment.CourseID)

	if err != nil {
		return err
	}
	enrollmentID, err := result.LastInsertId()
    if err != nil {
        return err
    }

    // Asignamos el ID a la entidad.
    enrollment.ID = int(enrollmentID)
    return nil

}