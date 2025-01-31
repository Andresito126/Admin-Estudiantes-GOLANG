package controllers

import (
	"net/http"
	"strconv"

	"github.com/Andresito126/go-estudiantes/src/courses/application"
	"github.com/Andresito126/go-estudiantes/src/courses/domain/entities"
	"github.com/gin-gonic/gin"
)

type UpdateCourseController struct{
	uc *application.UpdateCourseUseCase
}

func NewUpdateCourseController(uc *application.UpdateCourseUseCase) *UpdateCourseController{

	return &UpdateCourseController{uc: uc}
}


	func(control *UpdateCourseController)Run(c *gin.Context){
		var courses entities.Course

		if err := c.ShouldBindJSON(&courses); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalido"})
		return
	}
	if err := control.uc.Run(idInt, courses); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "La materia fue actualizada"})

	}

