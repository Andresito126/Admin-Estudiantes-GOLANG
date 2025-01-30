package main

import (
	

	"github.com/Andresito126/go-estudiantes/src/students/infrastructure/dependencies"
	"github.com/Andresito126/go-estudiantes/src/students/infrastructure/routes"
	"github.com/gin-gonic/gin"
)

func main() {
    dependencies.Init()
	r := gin.Default()
	routes.StudentRoutes(r)
	r.Run()
}