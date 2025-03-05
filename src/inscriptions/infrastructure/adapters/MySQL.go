package adapters

import (
	"database/sql"
	"fmt"

	"github.com/Andresito126/go-estudiantes/src/inscriptions/domain/entities"
	
)


type MySQL struct {
    db *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
    return &MySQL{db: db}
}

func (mysql *MySQL) Save(inscription *entities.Inscription) error {

    query := "INSERT INTO inscriptions (student_id, course_id, status) VALUES (?, ?, ?)"
    _, err := mysql.db.Exec(query, inscription.StudentID, inscription.CourseID, inscription.Status)

    if err != nil {
        return err
    }

    fmt.Println("Inscripción guardada con éxito")
    return nil
}


func (mysql *MySQL) FindByID(id int) (*entities.Inscription, error) {
    query := "SELECT * FROM inscriptions WHERE id = ?"

    row := mysql.db.QueryRow(query, id)
    var inscription entities.Inscription
    err := row.Scan(&inscription.ID, &inscription.StudentID, &inscription.CourseID, &inscription.Status)
    if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil 
		}
		return nil, err
    }
        fmt.Println("Inscripción encontrada")
        return &inscription, nil
}


func (mysql *MySQL) FindAll() ([]entities.Inscription, error) {
    query := "SELECT * FROM inscriptions"

    rows, err := mysql.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var inscriptions []entities.Inscription
    for rows.Next() {
        var inscription entities.Inscription
        if err := rows.Scan(&inscription.ID, &inscription.StudentID, &inscription.CourseID, &inscription.Status); err != nil {
            return nil, err
        }
        inscriptions = append(inscriptions, inscription)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return inscriptions, nil    
}


func (mysql *MySQL) Update(id int, inscription entities.Inscription) error {
    query := "UPDATE inscriptions SET student_id = ?, course_id = ?, status = ? WHERE id = ?"
    _, err := mysql.db.Exec(query, inscription.StudentID, inscription.CourseID, inscription.Status, id)
    fmt.Println("Inscripción actualizada")
    return err
}


