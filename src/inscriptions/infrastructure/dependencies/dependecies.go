package dependencies

import (
	"fmt"
	
	"github.com/Andresito126/go-estudiantes/src/core/database"
	"github.com/Andresito126/go-estudiantes/src/core/rabbitmq"
	"github.com/Andresito126/go-estudiantes/src/inscriptions/application"
	"github.com/Andresito126/go-estudiantes/src/inscriptions/application/services"
	"github.com/Andresito126/go-estudiantes/src/inscriptions/infrastructure/adapters"
	"github.com/Andresito126/go-estudiantes/src/inscriptions/infrastructure/controllers"
)

var (
	mySQL        adapters.MySQL
	rabbitRepo   *adapters.RabbitRepository
	eventService *services.EventService
)

func Init() {
	
	db, err := database.Connect()
	if err != nil {
		fmt.Println("Error en el servidor de la base de datos:", err)
		return
	}

	mySQL = *adapters.NewMySQL(db)

	
	rabbitConn, err := rabbitmq.Connect()
	if err != nil {
		fmt.Println("Error en la conexión a RabbitMQ:", err)
		return
	}
	//cone
	rabbitRepo = adapters.NewRabbitRepository(rabbitConn)

	//evento
	eventService = services.NewEventService(rabbitRepo)
}

func CreateInscriptionController() *controllers.SaveInscriptionController {
	ucCreateInscription := application.NewSaveInscriptionUseCase(&mySQL)
	return controllers.NewSaveInscriptionController(ucCreateInscription, eventService)
}

func FindAllInscriptionsController() *controllers.FindAllInscriptionsController {
	ucFindAllInscriptions := application.NewGetAllInscriptionsUseCase(&mySQL)
	return controllers.NewFindAllInscriptionsController(ucFindAllInscriptions)
}
