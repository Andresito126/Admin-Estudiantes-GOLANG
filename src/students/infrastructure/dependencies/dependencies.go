package dependencies

import (
	"fmt"

	"github.com/Andresito126/go-estudiantes/src/core/database"
	"github.com/Andresito126/go-estudiantes/src/students/application"
	"github.com/Andresito126/go-estudiantes/src/students/infrastructure/adapters"
	"github.com/Andresito126/go-estudiantes/src/students/infrastructure/controllers"
)

var (
	mySQl adapters.MySQL

)

func Init (){
	db, err := database.Connect()
	if err != nil{
		fmt.Println("Error en el server")
		return
	}

	mySQl = *adapters.NewMySQL(db)

}

//casos de uso

func CreateProductController() *controllers.CreateStudentController {
	ucCreateProduct := application.NewCreateStudentUseCase(&mySQl)

	return controllers.NewCreateStudentController(ucCreateProduct)
}