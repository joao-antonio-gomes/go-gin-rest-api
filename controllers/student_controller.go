package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joao-antonio-gomes/go-gin-rest-api/database"
	"github.com/joao-antonio-gomes/go-gin-rest-api/models"
)

func ShowAllStudents(ctx *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	ctx.JSON(http.StatusOK, &students)
}

func ShowStudent(ctx *gin.Context) {
	var student models.Student
	id := ctx.Params.ByName("id")
	database.DB.First(&student, id)

	if student.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Student not found!"})
		return
	}

	ctx.JSON(http.StatusOK, &student)
}

func CreateStudent(ctx *gin.Context) {
	var student models.Student

	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&student)
	ctx.JSON(http.StatusCreated, &student)
}

func DeleteStudent(ctx *gin.Context) {
	var student models.Student
	id := ctx.Params.ByName("id")
	database.DB.First(&student, id)

	if student.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Student not found!"})
		return
	}

	database.DB.Delete(&student, id)
	ctx.JSON(http.StatusNoContent, nil)
}

func EditStudent(ctx *gin.Context) {
	var student models.Student

	id := ctx.Params.ByName("id")
	database.DB.First(&student, id)

	if student.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Student not found!"})
		return
	}

	log.Println(&student)
	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	log.Println(&student)
	database.DB.Model(&student).UpdateColumns(student)
	ctx.JSON(http.StatusOK, &student)
}
