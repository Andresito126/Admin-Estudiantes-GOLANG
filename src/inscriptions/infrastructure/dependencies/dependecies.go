package dependencies

import (
	"fmt"

	"github.com/Andresito126/go-estudiantes/src/inscriptions/application"
	"github.com/Andresito126/go-estudiantes/src/inscriptions/infrastructure/controllers"
	"github.com/Andresito126/go-estudiantes/src/core/database"
	"github.com/Andresito126/go-estudiantes/src/inscriptions/infrastructure/adapters"
)

var (
	mySQL adapters.MySQL
)

func Init() {
	
	db, err := database.Connect()
	if err != nil {
		fmt.Println("Error en el servidor de la base de datos:", err)
		return
	}
	
	mySQL = *adapters.NewMySQL(db)
}



func CreateInscriptionController() *controllers.SaveInscriptionController {
	ucCreateInscription := application.NewSaveInscriptionUseCase(&mySQL)
	return controllers.NewSaveInscriptionController(ucCreateInscription)
}

func FindAllInscriptionsController() *controllers.FindAllInscriptionsController {
	ucFindAllInscriptions := application.NewGetAllInscriptionsUseCase(&mySQL)
	return controllers.NewFindAllInscriptionsController(ucFindAllInscriptions)
}

// func UpdateStatusController() *controllers.UpdateStatusController {
// 	ucUpdateStatus := application.NewUpdateStatusUseCase(&mySQL)
// 	return controllers.NewUpdateStatusController(ucUpdateStatus)
// }

// func GetByIdInscriptionController() *controllers.FindOneInscriptionController {
// 	ucGetByIdInscription := application.NewGetByIdInscription(&mySQL)
// 	return controllers.NewFindOneInscriptionController(ucGetByIdInscription)
// }


