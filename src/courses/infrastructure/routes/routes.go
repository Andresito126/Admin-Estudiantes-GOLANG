package routes

import (
	"github.com/Andresito126/go-estudiantes/src/courses/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

func CoursesRoutes(router *gin.Engine){
	routes := router.Group("/courses")

	saveCourse := dependencies.CreateCourseController()
	findAllCourses := dependencies.FindAllCourseController()
	updateCourse := dependencies.UpdateCourseController()
	deleteCourse := dependencies.DeleteCourseController()

	routes.POST("/", saveCourse.Run)
	routes.GET("/", findAllCourses.Run)
	routes.PUT("/:id", updateCourse.Run)
	routes.DELETE("/:id", deleteCourse.Run)

}