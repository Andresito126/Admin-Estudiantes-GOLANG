package routes

import (
	"github.com/Andresito126/go-estudiantes/src/students/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

func StudentRoutes(router *gin.Engine){

	routes := router.Group("students")

	saveStudent := dependencies.CreateProductController()
	finAllStudent := dependencies.FindAllStudentController()


	routes.POST("/add", saveStudent.Run)
	routes.GET("/", finAllStudent.Run)
}