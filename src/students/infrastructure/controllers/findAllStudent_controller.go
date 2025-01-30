package controllers

import (
	"net/http"

	"github.com/Andresito126/go-estudiantes/src/students/application"
	"github.com/gin-gonic/gin"
)

type FindAllStudentController struct{
	uc * application.FindAllStudentUseCase
}


func NewFindAllStudentController(uc *application.FindAllStudentUseCase)* FindAllStudentController{
	return &FindAllStudentController{uc:uc}
}

func (control *FindAllStudentController) Run(c *gin.Context) {
	students, err := control.uc.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"students": students})
}