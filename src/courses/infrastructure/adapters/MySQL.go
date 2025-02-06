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

func (mysql *MySQL) Save(course entities.Course) error {
	query := "INSERT INTO courses (name, duration) VALUES (?, ?)"
	_, err := mysql.db.Exec(query, course.Name, course.Duration)

	if err != nil {
		return err
	}

	fmt.Println("Materia guardado con éxito")
	return nil
}

func (mysql *MySQL) FindAll() ([]entities.Course, error) {
	query := "SELECT * FROM courses"

	rows, err := mysql.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []entities.Course
	for rows.Next() {
		var course entities.Course
		if err := rows.Scan(&course.ID, &course.Name, &course.Duration); err != nil {
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
	query := "UPDATE courses SET name = ?, duration = ? WHERE id = ?"
	_, err := mysql.db.Exec(query, course.Name,  course.Duration, id)
	fmt.Println("Materia actualizado")
	return err
}

func (mysql *MySQL) Delete(id int) error {
	query := "DELETE FROM courses WHERE id = ?"
	_, err := mysql.db.Exec(query, id)
	return err 
}
