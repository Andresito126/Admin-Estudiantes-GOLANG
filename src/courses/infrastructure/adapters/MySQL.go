package adapters

import (
	"database/sql"
	"fmt"

	"github.com/Andresito126/go-estudiantes/src/courses/domain/entities"
	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db: db}
}

func (mysql *MySQL) Save(course entities.Course) (int, error) {
	query := "INSERT INTO courses (name, duration, available_slots) VALUES (?, ?, ?)"
	result, err := mysql.db.Exec(query, course.Name, course.Duration, course.AvailableSlots)
	if err != nil {
		return 0, err
	}

	// Obtener el ID generado
	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	fmt.Println("Curso guardado con éxito con ID:", lastID)
	return int(lastID), nil
}



func (mysql *MySQL) FindAll() ([]entities.Course, error) {
	query := "SELECT id, name, duration, available_slots FROM courses"

	rows, err := mysql.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []entities.Course
	for rows.Next() {
		var course entities.Course
		if err := rows.Scan(&course.ID, &course.Name, &course.Duration, &course.AvailableSlots); err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return courses, nil
}

func (mysql *MySQL) Update(id int, course entities.Course) error {
	query := "UPDATE courses SET name = ?, duration = ?, available_slots = ? WHERE id = ?"
	_, err := mysql.db.Exec(query, course.Name, course.Duration, course.AvailableSlots, id)
	fmt.Println("Curso actualizado")
	return err
}

func (mysql *MySQL) Delete(id int) error {
	query := "DELETE FROM courses WHERE id = ?"
	_, err := mysql.db.Exec(query, id)
	return err 
}
