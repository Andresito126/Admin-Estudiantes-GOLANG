package main

import (
	

	studentsBD "github.com/Andresito126/go-estudiantes/src/students/infrastructure/dependencies"
	studentsRoutes "github.com/Andresito126/go-estudiantes/src/students/infrastructure/routes"
	coursesRoutes "github.com/Andresito126/go-estudiantes/src/courses/infrastructure/routes"
	coursesBD "github.com/Andresito126/go-estudiantes/src/courses/infrastructure/dependencies"
	enrollmentBD "github.com/Andresito126/go-estudiantes/src/enrollments/infrastructure/dependencies"
	enrollmentRoutes "github.com/Andresito126/go-estudiantes/src/enrollments/infrastructure/routes"
	"github.com/gin-gonic/gin"
)

func main() {
    studentsBD.Init()
	coursesBD.Init()
	enrollmentBD.Init()

	r := gin.Default()

	studentsRoutes.StudentRoutes(r)	
	coursesRoutes.CoursesRoutes(r)
	enrollmentRoutes.EnrollmentRoutes(r)
	r.Run()



}