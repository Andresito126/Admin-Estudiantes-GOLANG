package routes

import (
	"github.com/Andresito126/go-estudiantes/src/inscriptions/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

func InscriptionRoutes(router *gin.Engine) {

	
	routes := router.Group("inscriptions")
	
	saveInscription := dependencies.CreateInscriptionController()
	findAllInscriptions := dependencies.FindAllInscriptionsController()
	
	routes.POST("/", saveInscription.Run)                
	routes.GET("/", findAllInscriptions.Run)             
        
}
