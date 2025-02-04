package controllers

import (
	"net/http"

	"github.com/Andresito126/go-estudiantes/src/enrollments/application"
	"github.com/Andresito126/go-estudiantes/src/enrollments/domain/entities"
	"github.com/gin-gonic/gin"
)

type CreateEnrollmentController struct{
	uc * application.CreateEnrollmentUseCase
}

func NewCreateEnrollmentController(uc *application.CreateEnrollmentUseCase) *CreateEnrollmentController{
	return &CreateEnrollmentController{uc:uc}
}

func (control *CreateEnrollmentController) Run(ctx *gin.Context) {
    var enrollment entities.Enrollment

    
    if err := ctx.ShouldBindJSON(&enrollment); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos invalidos"})
        return
    }

    // se llama al use case para crear la inscripción.
    if err := control.uc.Run(&enrollment); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    // se creó.
    ctx.JSON(http.StatusCreated, enrollment)
}