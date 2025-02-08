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
	ucCreateStudent := application.NewCreateStudentUseCase(&mySQl)

	return controllers.NewCreateStudentController(ucCreateStudent)
}

func FindAllStudentController() *controllers.FindAllStudentController{
	ucFindAllStudent := application.NewFindAllStudentUseCase(&mySQl)
	
	return controllers.NewFindAllStudentController(ucFindAllStudent)
}

func UpdateStudentController() *controllers.UpdateStudentController {
	ucUpdateStudent := application.NewUpdateStudentUseCase(&mySQl)

	return controllers.NewUpdateStudentController(ucUpdateStudent)
}

func DeleteStudentController() *controllers.DeleteStudentController{
	ucDeleteStudent := application.NewDeleteStudentUseCase(&mySQl)
	
	return controllers.NewDeleteStudentController(ucDeleteStudent)
}


func FindByIdController() * controllers.FindByIdController{
	ucFindById := application.NewFindByIdUseCase(&mySQl)

	return controllers.NewFindByIdController(ucFindById)


}