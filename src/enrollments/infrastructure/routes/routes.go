package routes

import (
	"github.com/Andresito126/go-estudiantes/src/enrollments/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

func EnrollmentRoutes(router *gin.Engine){

	routes := router.Group("enrollments")

	createEnrollment := dependencies.CreateEnrollmentController()


	routes.POST("/", createEnrollment.Run)

	
}