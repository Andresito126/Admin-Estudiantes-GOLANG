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

    
    user := os.Getenv("RABBITMQ_USER")
    password := os.Getenv("RABBITMQ_PASSWORD")
    host := os.Getenv("RABBITMQ_HOST")
    port := os.Getenv("RABBITMQ_PORT")

    
    rabbitURL := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, host, port)

    //adptador
    rabbitRepo = adapters.NewRabbitRepository(rabbitURL)

    // event
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