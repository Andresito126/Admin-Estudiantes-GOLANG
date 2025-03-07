package controllers

import (
    "net/http"

    "github.com/Andresito126/go-estudiantes/src/inscriptions/application/services"
    application "github.com/Andresito126/go-estudiantes/src/inscriptions/application"
    "github.com/Andresito126/go-estudiantes/src/inscriptions/domain/entities"
    "github.com/gin-gonic/gin"
)

type SaveInscriptionController struct {
    useCase *application.SaveInscriptionUseCase
    event   *services.EventService
}

func NewSaveInscriptionController(useCase *application.SaveInscriptionUseCase, event *services.EventService) *SaveInscriptionController {
    return &SaveInscriptionController{
        useCase: useCase,
        event:   event,
    }
}

func (c *SaveInscriptionController) Run(ctx *gin.Context) {
    var inscription entities.Inscription

    // hace la conversion  
    if err := ctx.ShouldBindJSON(&inscription); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
        return
    }

    // validar campos obligatorios
    if inscription.StudentID <= 0 || inscription.CourseID <= 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Los campos student_id y course_id son obligatorios"})
        return
    }

    // guarda la inscripción usando el caso de uso
    if err := c.useCase.Run(&inscription); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo guardar la inscripción"})
        return
    }

    // publica el evento en reabbit en donde usa el servicio
    if err := c.event.PublishEvent(inscription); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo publicar el evento"})
        return
    }

    ctx.JSON(http.StatusCreated, gin.H{
        "message": "Inscripción en proceso ",
        "data":    inscription,
    })
}