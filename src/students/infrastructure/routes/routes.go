package routes

import (
	"github.com/Andresito126/go-estudiantes/src/students/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

func StudentRoutes(router *gin.Engine){

	routes := router.Group("students")

	saveStudent := dependencies.CreateProductController()
	finAllStudent := dependencies.FindAllStudentController()
	updateStudent := dependencies.UpdateStudentController()
	deleteStudent := dependencies.DeleteStudentController()
	findByIdStudent:= dependencies.FindByIdController()


	routes.POST("/", saveStudent.Run)
	routes.GET("/", finAllStudent.Run)
	routes.PUT("/:id", updateStudent.Run)
	routes.DELETE("/:id", deleteStudent.Run)
	routes.GET("/:id", findByIdStudent.Run )


}