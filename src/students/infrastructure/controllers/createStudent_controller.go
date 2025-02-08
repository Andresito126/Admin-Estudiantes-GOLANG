package controllers

import (
	"net/http"

	"github.com/Andresito126/go-estudiantes/src/students/application"
	"github.com/Andresito126/go-estudiantes/src/students/domain/entities"
	"github.com/gin-gonic/gin"
)

//onde implmentamos el use case
type CreateStudentController struct{
	uc * application.CreateStudentUseCase
}

//mi construct

func NewCreateStudentController(uc *application.CreateStudentUseCase) *CreateStudentController{

	return &CreateStudentController{uc:uc}

}

func (control *CreateStudentController) Run(c *gin.Context){

	var students entities.Student
	
	if err := c.ShouldBindJSON(&students); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Todos los campos son requeridos"})
		return 
	}

	err := control.uc.Run(students)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{"mensaje":"Estudiante agregado"})
		c.JSON(http.StatusOK, students)
	}

}


