package controllers

import (
	"net/http"
	"strconv"

	"github.com/Andresito126/go-estudiantes/src/students/application"
	"github.com/gin-gonic/gin"
)

type FindByIdController struct{
	uc *application.FindByIdUseCase
}


func NewFindByIdController (uc *application.FindByIdUseCase)* FindByIdController{

	return &FindByIdController{uc:uc}

}

func (control *FindByIdController)Run(c*gin.Context){

	id:= c.Param("id")
	studentID, err:= strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El id debe ser un numero"})
		return 
	}

	students, err := control.uc.Run(studentID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}


	c.JSON(http.StatusOK, gin.H{"students": students})
}