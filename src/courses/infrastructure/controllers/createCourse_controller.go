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
    var course entities.Course

    if err := c.ShouldBindJSON(&course); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de datos incorrecto"})
        return
    }

    //ejecuta y captura id
    id, err := control.uc.Run(course)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Asigna el ID obtenido (si el use case no lo hace internamente, hazlo aquí)
    course.ID = id

    c.JSON(http.StatusCreated, gin.H{
        "mensaje": "Materia agregada",
        "curso":   course,
    })
}

