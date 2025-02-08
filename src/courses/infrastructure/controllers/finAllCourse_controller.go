package controllers

import (
	"net/http"

	"github.com/Andresito126/go-estudiantes/src/courses/application"
	"github.com/gin-gonic/gin"
)

type FindAllCourseController struct{
	uc * application.FindAllCourseUseCase
}

func NewFindAllCourseController(uc *application.FindAllCourseUseCase)* FindAllCourseController{
	return &FindAllCourseController{uc:uc}
}

func(control *FindAllCourseController)Run(c* gin.Context){
	courses, err := control.uc.Run()
	if err!= nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"courses": courses})
}

