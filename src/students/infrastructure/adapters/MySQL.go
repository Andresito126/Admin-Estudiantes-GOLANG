package adapters

import (
	"database/sql"
	"fmt"

	"github.com/Andresito126/go-estudiantes/src/students/domain/entities"
	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db: db}
}

func (mysql *MySQL) Save(student entities.Student) error{
	
	query := "INSERT INTO students (name, email, career, matricula) VALUES (?, ?, ?, ?)"
	_, err := mysql.db. Exec(query,student.GetName(), student.GetEmail(), student.GetCareer(), student.GetMatricula())

	if err != nil {
		return err
	}

	fmt.Println("Estudiante guardado con éxito")
	return nil
}
