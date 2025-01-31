package controllers

import (
	"net/http"
	"strconv"

	"github.com/Andresito126/go-estudiantes/src/students/application"
	"github.com/Andresito126/go-estudiantes/src/students/domain/entities"
	"github.com/gin-gonic/gin"
)

type UpdateStudentController struct{
	uc *application.UpdateStudentUseCase
}

func NewUpdateStudentController(uc *application.UpdateStudentUseCase) *UpdateStudentController{

	return &UpdateStudentController{uc: uc}
}


	func(control *UpdateStudentController)Run(c *gin.Context){
		var students entities.Student

		if err := c.ShouldBindJSON(&students); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalido"})
		return
	}
	if err := control.uc.Run(idInt, students); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "El estudiante fue actualizado"})

	}

