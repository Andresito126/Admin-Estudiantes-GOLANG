package controllers

import (
	"net/http"

	"github.com/Andresito126/go-estudiantes/src/inscriptions/application"
	"github.com/gin-gonic/gin"
)


type FindAllInscriptionsController struct {
	uc *application.GetAllInscriptionsUseCase
}


func NewFindAllInscriptionsController(uc *application.GetAllInscriptionsUseCase) *FindAllInscriptionsController {
	return &FindAllInscriptionsController{uc: uc}
}

func (control *FindAllInscriptionsController) Run(c *gin.Context) {
	
	inscriptions, err := control.uc.Run()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"inscriptions": inscriptions})
}
