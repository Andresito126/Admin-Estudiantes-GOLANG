package dependencies

import (
	"fmt"

	"github.com/Andresito126/go-estudiantes/src/core/database"
	"github.com/Andresito126/go-estudiantes/src/courses/application"
	"github.com/Andresito126/go-estudiantes/src/courses/infrastructure/adapters"
	"github.com/Andresito126/go-estudiantes/src/courses/infrastructure/controllers"
)

var mySQl adapters.MySQL

func Init (){
	db, err := database.Connect()
	if err != nil{
		fmt.Println("Error en el server")
		return
	}

	mySQl = *adapters.NewMySQL(db)

}


func CreateCourseController() *controllers.CreateCourseController{
	ucCreateCourse:= application.NewSaveCourseUseCase(&mySQl)

	return controllers.NewCreateCourseController(ucCreateCourse)
}

func FindAllCourseController()* controllers.FindAllCourseController{
	ucFindAllCourse := application.NewFindAllCourseUseCase(&mySQl)
	
	return controllers.NewFindAllCourseController(ucFindAllCourse) 
}

func UpdateCourseController()*controllers.UpdateCourseController{
	ucUpdateCourse := application.NewUpdateCourseUseCase(&mySQl)

	return controllers.NewUpdateCourseController(ucUpdateCourse)
}

func DeleteCourseController() * controllers.DeleteCourseController{
	ucDeleteCourse := application.NewDeleteCourseUseCase(&mySQl)

	return controllers.NewDeleteCourseController(ucDeleteCourse)
}


