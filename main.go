package main

import (
	

	studentsBD "github.com/Andresito126/go-estudiantes/src/students/infrastructure/dependencies"
	studentsRoutes "github.com/Andresito126/go-estudiantes/src/students/infrastructure/routes"
	coursesRoutes "github.com/Andresito126/go-estudiantes/src/courses/infrastructure/routes"
	coursesBD "github.com/Andresito126/go-estudiantes/src/courses/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

func main() {
    studentsBD.Init()
	coursesBD.Init()
	r := gin.Default()

	studentsRoutes.StudentRoutes(r)	
	coursesRoutes.CoursesRoutes(r)
	r.Run()
}