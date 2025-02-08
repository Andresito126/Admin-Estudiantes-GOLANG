package controllers

import (
	"net/http"
	"strconv"

	"github.com/Andresito126/go-estudiantes/src/students/application"
	"github.com/gin-gonic/gin"
)

type DeleteStudentController struct {
	uc *application.DeleteStudentUseCase
}

func NewDeleteStudentController(uc *application.DeleteStudentUseCase)* DeleteStudentController{

	return &DeleteStudentController{uc:uc}
}

func (control *DeleteStudentController) Run(c *gin.Context) {
	id := c.Param("id")
	studentID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id de estudiante invalido"})
		return
	}

	err = control.uc.Run(studentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al encontrar el estudiante"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Estudiante eliminado "})

}