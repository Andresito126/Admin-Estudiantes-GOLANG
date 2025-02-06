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
	_, err := mysql.db. Exec(query, student.Name, student.Email, student.Career, student.Matricula)


	if err != nil {
		return err
	}

	fmt.Println("Estudiante guardado con éxito")
	return nil
}

func (mysql *MySQL) FindAll()([]entities.Student, error){
	query := "SELECT * FROM students"

rows, err := mysql.db.Query(query)
if err != nil {
	return nil, err
}
defer rows.Close()

var students []entities.Student
for rows.Next() {
	var student entities.Student
	if err := rows.Scan(&student.ID, &student.Name, &student.Email, &student.Career, &student.Matricula); err != nil {
		return nil, err
	}
	students = append(students, student)
}

if err := rows.Err(); err != nil {
	return nil, err
}

return students, nil
}

func (mysql *MySQL) Update(id int, student entities.Student) error {
	query := "UPDATE students SET name = ?, email= ?, career= ?, matricula = ? WHERE id = ?"
	_, err := mysql.db.Exec(query, student.Name, student.Email, student.Career, student.Matricula, id)
	fmt.Println("Estudiante actualizado")
	return err
}

func (mysql *MySQL) Delete(id int) error {
	query := "DELETE FROM students WHERE id = ?"
	_, err := mysql.db.Exec(query, id)
	return err 
}


func (mysql *MySQL) FindById(id int) ([]entities.Student, error) {
	query := "SELECT * FROM students WHERE id=?"
	rows, err := mysql.db.Query(query, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var students []entities.Student
	for rows.Next() {
		var student entities.Student
		err := rows.Scan(&student.ID, &student.Name, &student.Email, &student.Career, &student.Matricula)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	fmt.Println("Estudiante encontrado")
	return students, nil
}