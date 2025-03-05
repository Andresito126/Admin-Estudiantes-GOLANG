package controllers

import (
	"net/http"

	"github.com/Andresito126/go-estudiantes/src/inscriptions/application"
	"github.com/Andresito126/go-estudiantes/src/inscriptions/domain/entities"
	"github.com/gin-gonic/gin"
)


type SaveInscriptionController struct {
	uc *application.SaveInscriptionUseCase
}


func NewSaveInscriptionController(uc *application.SaveInscriptionUseCase) *SaveInscriptionController {
	return &SaveInscriptionController{uc: uc}
}

func (control *SaveInscriptionController) Run(c *gin.Context) {
	var inscription entities.Inscription

	
	if err := c.ShouldBindJSON(&inscription); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos los campos son requeridos"})
		return
	}

	
	err := control.uc.Run(&inscription)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{"mensaje": "Inscripción realizada con éxito"})
		c.JSON(http.StatusOK, inscription)
	}
}
