package controllers

import (
	"net/http"

	"github.com/Andresito126/go-estudiantes/src/courses/application"
	"github.com/Andresito126/go-estudiantes/src/courses/domain/entities"
	"github.com/gin-gonic/gin"
)

type CreateCourseController struct{
	uc * application.SaveCourseUseCase
}

func NewCreateCourseController(uc* application.SaveCourseUseCase)* CreateCourseController{
	return &CreateCourseController{uc:uc}
}

func (control *CreateCourseController) Run(c *gin.Context) {
	var courses entities.Course
	
	if err := c.ShouldBindJSON(&courses); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos los campos son requeridos"})
		return
	}

	err := control.uc.Run(courses)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"mensaje": "Materia agregada",
		"curso":   courses,
	})
}
