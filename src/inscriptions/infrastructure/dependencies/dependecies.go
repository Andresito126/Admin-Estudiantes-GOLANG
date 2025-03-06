package dependencies

import (
    "fmt"
    "os"

    "github.com/Andresito126/go-estudiantes/src/core/database"
    "github.com/Andresito126/go-estudiantes/src/inscriptions/application"
    "github.com/Andresito126/go-estudiantes/src/inscriptions/application/services"
    "github.com/Andresito126/go-estudiantes/src/inscriptions/infrastructure/adapters"
    "github.com/Andresito126/go-estudiantes/src/inscriptions/infrastructure/controllers"
)

var (
    mySQL        adapters.MySQL
    rabbitRepo   *adapters.RabbitRepository // adapter  de RabbitMQ
    eventService *services.EventService     // service de event
)

func Init() {
    
    db, err := database.Connect()
    if err != nil {
        fmt.Println("Error en el servidor de la base de datos:", err)
        return
    }
    mySQL = *adapters.NewMySQL(db)

    
    user := os.Getenv("RABBITMQ_USER")
    password := os.Getenv("RABBITMQ_PASSWORD")
    host := os.Getenv("RABBITMQ_HOST")
    port := os.Getenv("RABBITMQ_PORT")

    
    rabbitURL := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, host, port)

    //  adaptador de RabbitMQ
    rabbitRepo = adapters.NewRabbitRepository(rabbitURL)

    //  servicio de eventsew
    eventService = services.NewEventService(rabbitRepo)
}

func CreateInscriptionController() *controllers.SaveInscriptionController {
    ucCreateInscription := application.NewSaveInscriptionUseCase(&mySQL, eventService)
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


