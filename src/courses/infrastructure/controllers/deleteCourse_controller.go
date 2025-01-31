package controllers

import (
	"net/http"
	"strconv"

	"github.com/Andresito126/go-estudiantes/src/courses/application"
	"github.com/gin-gonic/gin"
)

type DeleteCourseController struct {
	uc *application.DeleteCourseUseCase
}

func NewDeleteCourseController(uc *application.DeleteCourseUseCase) *DeleteCourseController {

	return &DeleteCourseController{uc: uc}
}

func (control *DeleteCourseController) Run(c *gin.Context) {
	id := c.Param("id")
	courseID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id de la materia invalido"})
		return
	}

	err = control.uc.Run(courseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al encontrar la materia"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Materia eliminado "})

}
