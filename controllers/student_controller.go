package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/joao-antonio-gomes/go-gin-rest-api/models"
)

func ShowAllStudents(ctx *gin.Context) {
	ctx.JSON(200, models.Students)
}
