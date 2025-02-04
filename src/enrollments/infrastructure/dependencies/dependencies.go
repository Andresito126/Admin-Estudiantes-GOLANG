package dependencies

import (
	"fmt"

	"github.com/Andresito126/go-estudiantes/src/core/database"
	"github.com/Andresito126/go-estudiantes/src/enrollments/application"
	"github.com/Andresito126/go-estudiantes/src/enrollments/infrastructure/adapters"
	"github.com/Andresito126/go-estudiantes/src/enrollments/infrastructure/controllers"
)

var(mySQl adapters.MySQL)

func Init(){
	db, err := database.Connect()
	if err != nil{
		fmt.Println("Error server")
	}
	mySQl = *adapters.NewMySQL(db)
}


func CreateEnrollmentController() *controllers.CreateEnrollmentController{
	ucCreateEnrollment := application.NewCreateEnrollmentUseCase(&mySQl)

	return controllers.NewCreateEnrollmentController(ucCreateEnrollment)
}
