package main

import (
	coursesBD "github.com/Andresito126/go-estudiantes/src/courses/infrastructure/dependencies"
	coursesRoutes "github.com/Andresito126/go-estudiantes/src/courses/infrastructure/routes"
	studentsBD "github.com/Andresito126/go-estudiantes/src/students/infrastructure/dependencies"
	studentsRoutes "github.com/Andresito126/go-estudiantes/src/students/infrastructure/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
    studentsBD.Init()
	coursesBD.Init()
	r := gin.Default()
	r.Use(cors.Default())

	studentsRoutes.StudentRoutes(r)	
	coursesRoutes.CoursesRoutes(r)
	r.Run()
}